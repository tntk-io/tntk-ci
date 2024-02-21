import axios from "axios";
import { useSelector } from "react-redux"
import { Redirect } from "react-router-dom";
import { Route, Switch } from "react-router-dom";
import { Router } from "react-router-dom/cjs/react-router-dom.min";
import { history } from "../../helpers/api";
import Home from "../../pages/Home";
import Login from "../../pages/login";
import Register from "../../pages/register";
import PrivateRoute from "./private-route";

function AppNavigator() {
    const Error = () => <div>Not found</div>

    const user = useSelector((state) => state.auth);

    if (user && user.token) {
        axios.defaults.headers.common["Authorization"] = `Token ${user.token}`;
    }
    const loggedIn = useSelector((state) => state.auth.loggedIn);

    return (
        <Router history={history}>
            <Switch>
                <Route path="/login">
                    {loggedIn ? <Redirect to="/" /> : <Login />}
                </Route>
                <Route path="/register">
                    <Register />
                </Route>
                {user && user.user ? <PrivateRoute exact path="/" component={Home} user={user} /> : <Login />}                
                <PrivateRoute path="*" component={Error} user={user} />
            </Switch>
        </Router>
    )
}

export default AppNavigator;
