import React, { useEffect, useState } from 'react';
import { Flex } from '@patternfly/react-core';

import { SearchFilter } from 'types/search';
import {
    OnSearchPayload,
    PartialCompoundSearchFilterConfig,
    SearchFilterAttributeName,
    SearchFilterEntityName,
} from '../types';
import { getDefaultAttribute, getDefaultEntity } from '../utils/utils';

import EntitySelector, { SelectedEntity } from './EntitySelector';
import AttributeSelector, { SelectedAttribute } from './AttributeSelector';
import CompoundSearchFilterInputField, { InputFieldValue } from './CompoundSearchFilterInputField';

export type CompoundSearchFilterProps = {
    config: PartialCompoundSearchFilterConfig;
    defaultEntity?: SearchFilterEntityName;
    defaultAttribute?: SearchFilterAttributeName;
    searchFilter: SearchFilter;
    additionalContextFilter?: SearchFilter;
    onSearch: ({ action, category, value }: OnSearchPayload) => void;
};

function CompoundSearchFilter({
    config,
    defaultEntity,
    defaultAttribute,
    searchFilter,
    additionalContextFilter,
    onSearch,
}: CompoundSearchFilterProps) {
    const [selectedEntity, setSelectedEntity] = useState<SelectedEntity>(() => {
        if (defaultEntity) {
            return defaultEntity;
        }
        return getDefaultEntity(config);
    });

    const [selectedAttribute, setSelectedAttribute] = useState<SelectedAttribute>(() => {
        if (defaultAttribute) {
            return defaultAttribute;
        }
        const defaultEntity = getDefaultEntity(config);
        if (!defaultEntity) {
            return undefined;
        }
        return getDefaultAttribute(defaultEntity, config);
    });

    const [inputValue, setInputValue] = useState<InputFieldValue>('');

    useEffect(() => {
        if (defaultEntity) {
            setSelectedEntity(defaultEntity);
        }
    }, [defaultEntity]);

    useEffect(() => {
        if (defaultAttribute) {
            setSelectedAttribute(defaultAttribute);
        }
    }, [defaultAttribute]);

    return (
        <Flex
            direction={{ default: 'row' }}
            spaceItems={{ default: 'spaceItemsNone' }}
            flexWrap={{ default: 'nowrap' }}
            className="pf-v5-u-w-100"
        >
            <EntitySelector
                menuToggleClassName="pf-v5-u-flex-shrink-0"
                selectedEntity={selectedEntity}
                onChange={(value) => {
                    setSelectedEntity(value as SearchFilterEntityName);
                    const defaultAttribute = getDefaultAttribute(
                        value as SearchFilterEntityName,
                        config
                    );
                    setSelectedAttribute(defaultAttribute);
                    setInputValue('');
                }}
                config={config}
            />
            <AttributeSelector
                menuToggleClassName="pf-v5-u-flex-shrink-0"
                selectedEntity={selectedEntity}
                selectedAttribute={selectedAttribute}
                onChange={(value) => {
                    setSelectedAttribute(value as SearchFilterAttributeName);
                    setInputValue('');
                }}
                config={config}
            />
            <CompoundSearchFilterInputField
                selectedEntity={selectedEntity}
                selectedAttribute={selectedAttribute}
                value={inputValue}
                onChange={(value) => {
                    setInputValue(value);
                }}
                searchFilter={searchFilter}
                additionalContextFilter={additionalContextFilter}
                onSearch={(payload) => {
                    const { action, category, value } = payload;
                    const shouldSearch =
                        (action === 'ADD' &&
                            value !== '' &&
                            !searchFilter?.[category]?.includes(value)) ||
                        (action === 'REMOVE' && value !== '');

                    if (shouldSearch) {
                        onSearch(payload);
                    }
                }}
                config={config}
            />
        </Flex>
    );
}

export default CompoundSearchFilter;
