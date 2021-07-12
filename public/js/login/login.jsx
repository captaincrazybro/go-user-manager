class Page extends React.Component {
    constructor() {
        super();
        this.state = {
            session: getCookie("session"),
            user: getCookie("user"),
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
            username: "",
            password: ""
        }

        this.handleUsernameChange = this.handleUsernameChange.bind(this)
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
    render() {
        return (
            <div id={"login-form"}>
                <form action={"/api/login"} method={"post"}>
                    <div id={"login-form-fields"}>
                        <label class={"login-form-label"} for={"login-username"}>Username</label><br />
                        <input id={"login-username"} class={"login-form-field"} type={"text"} value={this.state.username} onChange={} placeholder={"Enter username"} required/><br />
                        <label class={"login-form-label"} for={"login-password"}>Password</label><br />
                        <input id={"login-password"} class={"login-form-field"} type={"text"} value={this.state.password} placeholder={"Enter password"} required/><br /><br />
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