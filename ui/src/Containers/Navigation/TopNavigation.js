import React, { Component } from 'react';
import PropTypes from 'prop-types';
import ReactRouterPropTypes from 'react-router-prop-types';
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux';
import { actions as globalSearchActions } from 'reducers/globalSearch';
import { createStructuredSelector } from 'reselect';
import { selectors } from 'reducers';

import * as Icon from 'react-feather';
import Logo from 'Components/icons/logo';
import AuthService from 'services/AuthService';

const titleMap = {
    numAlerts: 'Violation',
    numClusters: 'Cluster',
    numDeployments: 'Deployment',
    numImages: 'Image'
};

class TopNavigation extends Component {
    static propTypes = {
        history: ReactRouterPropTypes.history.isRequired,
        toggleGlobalSearchView: PropTypes.func.isRequired,
        summaryCounts: PropTypes.shape({
            numAlerts: PropTypes.string,
            numClusters: PropTypes.string,
            numDeployments: PropTypes.string,
            numImages: PropTypes.string
        })
    };

    static defaultProps = {
        summaryCounts: null
    };

    renderLogoutButton = () => {
        if (!AuthService.isLoggedIn()) return '';
        const logout = () => () => {
            AuthService.logout();
            this.props.history.push('/login');
        };
        return (
            <button
                onClick={logout()}
                className="flex flex-end border-l border-r border-base-300 px-4 no-underline py-3 text-base-600 hover:bg-base-200 items-center cursor-pointer"
            >
                <Icon.LogOut className="h-4 w-4 mr-3" />
                <span className="uppercase text-sm tracking-wide">Logout</span>
            </button>
        );
    };

    renderSearchButton = () => (
        <button
            onClick={this.props.toggleGlobalSearchView}
            className="flex flex-end border-l border-r border-base-300 px-4 no-underline py-3 text-base-600 hover:bg-base-200 items-center cursor-pointer"
        >
            <Icon.Search className="h-4 w-4 mr-3" />
            <span className="uppercase text-sm tracking-wide">Search</span>
        </button>
    );

    renderSummaryCounts = () => {
        const { summaryCounts } = this.props;
        if (!summaryCounts) return '';
        return (
            <ul className="flex uppercase text-sm p-0 w-full">
                {Object.keys(summaryCounts).map(key => (
                    <li
                        key={key}
                        className="flex flex-1 flex-col border-r border-base-300 px-4 no-underline py-3 text-base-500 items-center"
                    >
                        <div className="text-xl">{summaryCounts[key]}</div>
                        <div className="text-sm pt-1">
                            {titleMap[key]}
                            {summaryCounts[key] === '1' ? '' : 's'}
                        </div>
                    </li>
                ))}
            </ul>
        );
    };

    render() {
        return (
            <nav className="top-navigation flex flex-row flex-1 justify-between bg-base-100 border-base-300 border-b relative">
                <div className="flex w-full max-w-sm">
                    <div className="py-2 px-5 border-r bg-white border-base-400">
                        <Logo className="fill-current text-primary-800 h-10 w-10 " />
                    </div>
                    {this.renderSummaryCounts()}
                </div>
                <div className="flex pr-4">
                    {this.renderSearchButton()}
                    {this.renderLogoutButton()}
                </div>
            </nav>
        );
    }
}

const mapStateToProps = createStructuredSelector({
    summaryCounts: selectors.getSummaryCounts
});

const mapDispatchToProps = dispatch => ({
    toggleGlobalSearchView: () => dispatch(globalSearchActions.toggleGlobalSearchView())
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(TopNavigation));
