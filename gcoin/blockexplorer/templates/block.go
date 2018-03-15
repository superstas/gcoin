package templates

// BlockTPL represents a body of the block page
const BlockTPL = `{{template "header"}}
<div class="row col-12">
        <span class="btn btn-primary mx-auto">Block hash: {{.Block.Hash}}</span>
    </div>
    <hr class="col-xs-12">
    <div class="row col-12">
        <div class="mx-auto">
        <h4>Header</h4>
        <table class="table table-bordered">
            <thead>
            <tr>
                <th scope="col">Field</th>
                <th scope="col">Value</th>
            </tr>
            </thead>
            <tbody>
            <tr >
                <th scope="row">PreviousBlockHash</th>
                <td>
					{{ if .Block.Header.PreviousBlockHash }}
					<a href="/block/{{.Block.Header.PreviousBlockHash}}" class="btn btn-outline-dark">{{.Block.Header.PreviousBlockHash}}</a>
					{{ else }}
					<button class="btn btn-outline-success">Genesis block</button>
					{{ end }}
				</td>
            </tr>
            <tr>
                <th scope="row">MerkleRootHash</th>
                <td>{{.Block.Header.MerkleRootHash}}</td>
            </tr>
            <tr>
                <th scope="row">Timestamp</th>
                <td>{{.Block.Header.Timestamp}}</td>
            </tr>
            <tr>
                <th scope="row">Target</th>
                <td>{{.Block.Header.Target}}</td>
            </tr>
            <tr>
                <th scope="row">Nonce</th>
                <td>{{.Block.Header.Nonce}}</td>
            </tr>
            </tbody>
        </table>
        </div>
    </div>
    <div class="row col-12">
        <div class="mx-auto">
			{{ if .Block.Transactions }}
	            <h4>Transactions</h4>
				{{range $i, $tx := .Block.Transactions}}
				<p>
                	<a href="/tx/{{$tx.ID}}" class="btn btn-outline-primary">{{$tx.ID}}</a>
            	</p>
				{{ end }}
			{{ end }}
        </div>
    </div>
    <hr class="col-xs-12">
    <div class="row col-12">
        <div class="mx-auto">
            <h5><a class="btn btn-info" data-toggle="collapse" href="#rawJSON" role="button">Show raw JSON</a></h5>
            <div class="collapse" id="rawJSON">
                <pre>{{ .RawJSON }}</pre>
            </div>
        </div>
    </div>{{template "footer"}}
`
