import React, { useContext } from 'react';
import ReactRouterPropTypes from 'react-router-prop-types';
import entityLabels from 'messages/entity';
import pluralize from 'pluralize';
import URLService from 'modules/URLService';
import startCase from 'lodash/startCase';

import SidePanelAnimation from 'Components/animations/SidePanelAnimation';

import PageHeader from 'Components/PageHeader';
import ExportButton from 'Components/ExportButton';
import searchContext from 'Containers/searchContext';
import { searchParams } from 'constants/searchParams';
import List from './EntityList';
import SidePanel from '../SidePanel/SidePanel';

const ListPage = ({ match, location, history }) => {
    const params = URLService.getParams(match, location);
    const {
        pageEntityListType,
        entityId1,
        entityType2,
        entityListType2,
        entityId2,
        query
    } = params;
    const searchParam = useContext(searchContext);

    function onRowClick(entityId) {
        const urlBuilder = URLService.getURL(match, location).push(entityId);
        history.push(urlBuilder.url());
    }

    const header = pluralize(entityLabels[pageEntityListType]);
    const exportFilename = `${pluralize(startCase(header))} Report`;
    return (
        <>
            <PageHeader header={header} subHeader="Entity List">
                <div className="flex flex-1 justify-end">
                    <div className="flex">
                        <div className="flex items-center">
                            <ExportButton
                                fileName={exportFilename}
                                type={pageEntityListType}
                                page="configManagement"
                                pdfId="capture-list"
                            />
                        </div>
                    </div>
                </div>
            </PageHeader>
            <div className="flex flex-1 h-full relative z-0">
                <List
                    className={entityId1 ? 'overlay' : ''}
                    entityListType={pageEntityListType}
                    entityId={entityId1}
                    onRowClick={onRowClick}
                    query={query[searchParam]}
                />
                <searchContext.Provider value={searchParams.sidePanel}>
                    <SidePanelAnimation condition={!!entityId1}>
                        <SidePanel
                            className="w-full h-full bg-base-100 border-l border-base-400 shadow-sidepanel"
                            entityType1={pageEntityListType}
                            entityId1={entityId1}
                            entityType2={entityType2}
                            entityListType2={entityListType2}
                            entityId2={entityId2}
                            query={query}
                        />
                    </SidePanelAnimation>
                </searchContext.Provider>
            </div>
        </>
    );
};

ListPage.propTypes = {
    match: ReactRouterPropTypes.match.isRequired,
    location: ReactRouterPropTypes.location.isRequired,
    history: ReactRouterPropTypes.history.isRequired
};

export default ListPage;
