class Welcome extends React.Component {
    render() {
        // fetch("http://localhost:8080/api/ping")
        //     .then(val => val.text())
        //     .then(val => console.log(val))
        return (
            <div id="welcome" class="element">
                <div style={{margin: "5px"}}>
                    <strong>
                        <a id="welcome-link" href="/" style={{textDecoration: "none"}}>Welcome to the GTIS Mack User Manager!</a>
                    </strong>
                </div>
            </div>
        )
    }
}