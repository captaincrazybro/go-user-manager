import React from 'react'
import ReactDOM from 'react-dom'

class Page extends React.Component {
    render() {
        return (
            <div>
                <h1>Hello world!</h1>
                <p>How are you doing?</p>
            </div>
        )
    }

}

class Home extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>Jokeish</h1>
                    <p>A load of Dad jokes XD</p>
                    <p>Sign in to get access </p>
                    <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
                </div>
            </div>
        )
    }
}

ReactDOM.render(<Page />, document.getElementById("page"))