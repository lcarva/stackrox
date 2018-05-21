package blevesearch

import (
	"reflect"
	"time"

	"bitbucket.org/stack-rox/apollo/central/metrics"
	"bitbucket.org/stack-rox/apollo/central/search"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/images"
	"github.com/blevesearch/bleve"
	"github.com/deckarep/golang-set"
)

// maxDeploymentsReturned is for the specific scoped request to deployments which is then joined with images
// so we need to return as many deployments as possible to not exclude relevant images
const maxDeploymentsReturned = 500

var imageObjectMap = map[string]string{
	"image": "",
}

// AddImage adds the image to the index
func (b *Indexer) AddImage(image *v1.Image) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), "Add", "Image")
	digest := images.NewDigest(image.GetName().GetSha()).Digest()
	return b.imageIndex.Index(digest, image)
}

// DeleteImage deletes the image from the index
func (b *Indexer) DeleteImage(sha string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), "Delete", "Image")
	return b.imageIndex.Delete(sha)
}

// SearchImages takes a SearchRequest and finds any matches
// It is different from other requests because it requires that we actually search the deployments
// for the image that may match the criteria
func (b *Indexer) SearchImages(request *v1.ParsedSearchRequest) (results []search.Result, err error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), "Search", "Image")
	// If we have scopes set, or the request has nothing set, get the list of image SHAs from the deployments.
	// If no scopes are set, this returns ALL image SHAs present in deployments.
	var imageSHAs mapset.Set
	if len(request.GetScopes()) > 0 || (len(request.GetFields()) == 0 && request.GetStringQuery() == "") {
		if imageSHAs, err = b.getImageSHAsFromScopes(request.GetScopes()); err != nil {
			return
		}
	}

	// If there is not query or query files, then we don't need to filter our image set.
	if len(request.GetFields()) == 0 && request.GetStringQuery() == "" {
		results, err = shasToResults(imageSHAs)
		return
	}

	// Create and run query for fields, and input string query, if it exists.
	imageQuery, err := fieldsToQuery(request.GetFields(), imageObjectMap)
	if err != nil {
		return nil, err
	}
	if request.GetStringQuery() != "" {
		imageQuery.AddQuery(bleve.NewQueryStringQuery(request.GetStringQuery()))
	}
	results, err = runQuery(imageQuery, b.imageIndex)
	if err != nil {
		return nil, err
	}

	// Filter results by which fields exist in the results retrieved from the deployments.
	if imageSHAs != nil {
		filteredResults := results[:0]
		for _, result := range results {
			if imageSHAs.Contains(result.ID) {
				filteredResults = append(filteredResults, result)
			}
		}
		results = filteredResults
	}
	return
}

func (b *Indexer) getImageSHAsFromScopes(scopes []*v1.Scope) (mapset.Set, error) {
	scopesQuery := getScopesQuery(scopes, scopeToDeploymentQuery)
	searchRequest := bleve.NewSearchRequest(scopesQuery)
	searchRequest.Fields = []string{"containers.image.name.sha"}
	searchRequest.Size = maxDeploymentsReturned

	searchResult, err := b.deploymentIndex.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	return deploymentResultsToShaSet(searchResult)
}

func shasToResults(shas mapset.Set) ([]search.Result, error) {
	if shas == nil {
		return nil, nil
	}

	searchResults := make([]search.Result, 0, shas.Cardinality())
	for sha := range shas.Iter() {
		searchResults = append(searchResults, search.Result{ID: sha.(string)})
	}
	return searchResults, nil
}

func deploymentResultsToShaSet(searchResult *bleve.SearchResult) (mapset.Set, error) {
	shaSetFromDeployment := mapset.NewSet()
	for _, hit := range searchResult.Hits {
		shaObj := hit.Fields["containers.image.name.sha"]
		if shaObj == nil {
			continue
		}
		t := reflect.TypeOf(shaObj)
		kind := t.Kind()
		switch kind {
		case reflect.Slice:
			strSlice, ok := shaObj.([]interface{})
			if !ok {
				logger.Errorf("Unexpected Slice type %s for image sha", t)
				continue
			}
			for _, s := range strSlice {
				shaSetFromDeployment.Add(s.(string))
			}
		case reflect.String:
			shaSetFromDeployment.Add(shaObj.(string))
		default:
			logger.Errorf("Unexpected type %s for image sha", t)
		}
	}
	return shaSetFromDeployment, nil
}
