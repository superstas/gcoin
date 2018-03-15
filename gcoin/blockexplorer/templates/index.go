package templates

// IndexTPL represents a body of the main page
const IndexTPL = `{{template "header"}}
<div class="row col-12">
    <div class="mx-auto">
        <h4>Last 10 blocks</h4>
        {{range $block := .Blocks}}
        <p>
            {{if $block.PreviousBlockHash }}
                <a href="/block/{{$block.HexHash}}" class="btn btn-outline-dark">{{$block.HexHash}}</a>
            {{else}}
                <a href="/block/{{$block.HexHash}}" class="btn btn-outline-primary" title="Genesis block">{{$block.HexHash}}</a>
            {{end}}
        </p>
        {{ end }}
    </div>
</div>
<hr class="col-xs-12">
<div class="row col-12">
    <div class="mx-auto">
		{{ if .Transactions }}
        <h4>Unconfirmed transactions</h4>
	        {{range $tx := .Transactions}}
    	    <p>
        	    <a href="/tx/{{$tx.HexID}}" class="btn btn-outline-danger">{{$tx.HexID}}</a>
        	</p>
        	{{ end }}
        {{ end }}
    </div>
</div>
{{template "footer"}}`
