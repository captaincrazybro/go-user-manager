class Page extends React.Component {
    constructor() {
        super();
    }

    render() {
        return (
            "hi"
        )
    }
}

ReactDOM.render((<div>
    <Top />
    <div id="page">
        <Page />
    </div>
</div>), document.getElementById("root"))