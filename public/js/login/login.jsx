class Page extends React.Component {
    constructor() {
        super();
    }
}

ReactDOM.render((<div>
    <div id="top">
        <Welcome /><Login />
    </div>
    <div id="page">
        <Page />
    </div>
</div>), document.getElementById("root"))