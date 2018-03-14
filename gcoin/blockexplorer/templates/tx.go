package templates

const TXTpl = `{{template "header"}}
    <div class="row col-12">
        <div class="mx-auto">
			{{ if .Confirmed }}
			    <span class="btn btn-success">Transaction ID: {{.TX.ID}}</span>
            	<span>included in <a href="/block/{{.BlockHash}}" class="btn btn-outline-dark">block</a></span>
        	
			{{ else }}
				<span class="btn btn-danger">Transaction ID: {{.TX.ID}}</span>
			{{ end }}
        </div>
    </div>
    <hr class="col-xs-12">
    <div class="row col-12">
        <h4>Inputs</h4>
        <table class="table table-bordered">
            <thead>
            <tr>
                <th scope="col">TxID</th>
                <th scope="col">OutIndex</th>
                <th scope="col">Sign</th>
                <th scope="col">PubKey</th>
            </tr>
            </thead>
            <tbody>
			{{range $i, $input := .TX.Inputs}}
			<tr>
                <th scope="row">
					{{ if $input.TransactionID }}	
						<a href="/tx/{{$input.TransactionID}}" class="btn btn-outline-primary">{{$input.TransactionID}}</a></th>
					{{ else }}
						<button class="btn btn-outline-success">Coinbase TX</button>
					{{ end }}
                <td>{{$input.OutIndex}}</td>
                <td>
                    <a class="btn btn-info" data-toggle="collapse" href="#sign_{{$i}}" role="button">Show sign</a>
                    <div class="collapse" id="sign_{{$i}}">
                        <br>
                        <div class="card card-body">{{$input.Sign}}</div>
                    </div>

                </td>
                <td>
                    <a class="btn btn-info" data-toggle="collapse" href="#pubkey_{{$i}}" role="button">Show pubkey</a>
                    <div class="collapse" id="pubkey_{{$i}}">
                        <br>
                        <div class="card card-body">{{$input.PubKey}}</div>
                    </div>
                </td>
            </tr>
	        {{ end }}
            </tbody>
        </table>
    </div>
   <hr class="col-xs-12">
   <div class="row col-12">
        <h4>Outputs</h4>
        <table class="table table-bordered">
            <thead>
            <tr>
                <th scope="col">Address</th>
                <th scope="col">Value</th>
            </tr>
            </thead>
            <tbody>
            {{range $i, $output := .TX.Outputs}}
				<tr>
                	<th scope="row">{{$output.Address}}</th>
                	<td>{{$output.Value.String}}</td>
            	</tr>
			{{end}}
            </tbody>
        </table>
    </div>
    <div class="row col-12">
        <div>
        <h5><a class="btn btn-info" data-toggle="collapse" href="#rawJSON" role="button">Show raw JSON</a></h5>
        <div class="collapse" id="rawJSON">
        <pre>{{ .RawJSON }}</pre>
        </div>
    </div>
</div>{{template "footer"}}`
