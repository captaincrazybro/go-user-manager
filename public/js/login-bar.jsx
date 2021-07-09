class Login extends React.Component {
    constructor() {
        super();
        this.state = {
            session: localStorage.getItem("session"),
            user: localStorage.getItem("user"),
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
            <div id="login" className="element">
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
            <div>
                <strong>Hello {this.props.user}!</strong>
                <button id="login-button" onClick={this.handleClick}>Logout</button>
            </div>
        )
    }
}

class NotLoggedIn extends React.Component {
    constructor() {
        super();
        this.handleClick = this.handleClick.bind(this)
    }

    handleClick() {
        localStorage.setItem("session", "asdf1234")
        localStorage.setItem("user", "cqptain")
        window.location.reload(false);
    }

    render() {
        return (
            <div>
                <strong>You are not logged in!</strong>
                <button id="login-button" onClick={this.handleClick} /*href="/login"*/>Login</button>
            </div>
        )
    }
}