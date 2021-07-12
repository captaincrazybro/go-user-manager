class Page extends React.Component {
    constructor() {
        super();
        this.state = {
            session: localStorage.getItem("session"),
            user: localStorage.getItem("user"),
        }
        this.state.loggedIn = this.state.session && this.state.user
    }

    render() {
        let pageContent;
        if (this.state.loggedIn) {
            pageContent = (
                <h3>You are already logged in</h3>
            )
        } else {
            pageContent = <LoginForm />
        }

        return (
            <div id="login-page" style={{margin: "0 auto"}} class={"page-element"}>
                {pageContent}
            </div>
        )
    }
}

class LoginForm extends React.Component {
    constructor() {
        super();
        this.state = {
            error: null,
            username: "",
            password: ""
        }

        this.handleUsernameChange = this.handleUsernameChange.bind(this)
        this.handlePasswordChange = this.handlePasswordChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    handleUsernameChange(event) {
        this.setState({
            username: event.target.value
        })
    }

    handlePasswordChange(event) {
        this.setState({
            password: event.target.value
        })
    }

    async handleSubmit(event) {
        event.preventDefault()
        let username = this.state.username
        let res = await fetch("/api/login", {method: "POST", body: JSON.stringify({username: username, password: this.state.password}), headers: {'Accept': 'application/json', 'Content-Type': 'application/json'}})
        let json = await res.json()

        if (json.type == "error") this.setState({error: json.message}); return

        let session = json.json.sessionId

        localStorage.setItem("session", session)
        localStorage.setItem("user", username)
        window.location.reload(false);
    }
    render() {
        let error
        if (this.state.error) {
            error = (
                <h5 style={{color: "red", textAlign: "left"}}>{this.state.error}</h5>
            )
        }

        return (
            <div id={"login-form"}>
                {error}
                <form onSubmit={this.handleSubmit}>
                    <div id={"login-form-fields"}>
                        <label class={"login-form-label"} for={"login-username"}>Username</label><br />
                        <input id={"login-username"} class={"login-form-field"} type={"text"} value={this.state.username} onChange={this.handleUsernameChange} placeholder={"Enter username"} required/><br />
                        <label class={"login-form-label"} for={"login-password"}>Password</label><br />
                        <input id={"login-password"} class={"login-form-field"} type={"password"} value={this.state.password} onChange={this.handlePasswordChange} placeholder={"Enter password"} required/><br /><br />
                    </div>
                    <div id={"login-form-button"}><input id={"login-button"} type={"submit"} value={"Sign In"} /></div>
                </form>
            </div>
        )
    }
}

ReactDOM.render((<div>
    <Top />
    <br />
    <br />
    <br />
    <div id="page">
        <Page />
    </div>
</div>), document.getElementById("root"))