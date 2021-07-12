class Top extends React.Component {
    render() {
        return (
            <div id="top" className="page-element">
                    <Welcome/> <Login/>
            </div>
        )
    }
}

class Login extends React.Component {
    constructor() {
        super();
        this.state = {
            session: localStorage.getItem("session"),
            user: localStorage.getItem("user")
        }
        this.state.loggedIn = this.state.session && this.state.user
    }
    render() {
        let login;
        if (this.state.loggedIn) {
            login = <LoggedIn  user={this.state.user}/>
        } else {
            login = <NotLoggedIn />
        }

        return (
            <div id="login">
                {login}
            </div>
        )
    }
}

class LoggedIn extends React.Component {
    constructor(props) {
        super(props);
        this.handleClick = this.handleClick.bind(this)
    }

    handleClick() {
        // TODO: post to /api/logout
        localStorage.removeItem("session")
        localStorage.removeItem("user")
        window.location.reload(false);
    }
    render() {
        return (
            <div style={{margin: "5px"}}>
                <strong>Hello {this.props.user}!</strong>
                <button id="login-button" onClick={this.handleClick}>Logout</button>
            </div>
        )
    }
}

class NotLoggedIn extends React.Component {
    constructor() {
        super();
        // this.handleClick = this.handleClick.bind(this)
    }

    // handleClick() {
    //     localStorage.setItem("session", "asdf1234")
    //     localStorage.setItem("user", "cqptain")
    //     window.location.reload(false);
    // }

    render() {
        return (
            <div style={{margin: "5px"}}>
                <strong>You are not logged in!</strong>
                <a href="/login">
                    <button id="login-button">Login</button>
                </a>
            </div>
        )
    }
}

class Welcome extends React.Component {
    render() {
        return (
            <div id="welcome" style={{margin: "5px"}}>
                    <strong>
                        <a id="welcome-link" href="/" style={{textDecoration: "none"}}>Welcome to the GTIS Mack User Manager!</a>
                    </strong>
            </div>
        )
    }
}

/**
 *
 * @deprecated
 * @param cname
 * @param cvalue
 */
function setCookie(cname, cvalue) {
    //const d = new Date();
    //d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
    //let expires = "expires="+d.toUTCString();
    document.cookie = cname + "=" + cvalue +  + ";path=/";
}

/**
 *
 * @deprecated
 * @param cname
 * @returns {string}
 */
function getCookie(cname) {
    let name = cname + "=";
    let ca = document.cookie.split(';');
    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) === ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) === 0) {
            return c.substring(name.length, c.length);
        }
    }
}

/**
 *
 * @deprecated
 * @param cname
 */
function deleteCookie(cname) {
    document.cookie = `${cname}=${getCookie(cname)}; max-age=0;`
}