import React from 'react'
import history from '../history'
import Login from '../pages/Auth'
import {Switch, Router, Route} from 'react-router-dom'
import {Dashboard} from '../pages';

const MainRoutes = () => (
    <Router history={history}>
        <Switch>
            <Route exact path="/" component={Dashboard}/>
            <Route exact path="/login" component={Login}/>
        </Switch>
    </Router>
);
export default MainRoutes;