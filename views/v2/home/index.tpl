{{template "v2/base.tpl" .}}

{{define "optional_scripts_block"}}
<script src="{{cdnjs "/static/adminlte/plugins/chart.js/Chart.min.js"}}"></script>
<!-- AdminLTE for demo purposes -->
<script src="{{cdnjs "/static/adminlte//js/demo.js"}}"></script>
<!-- AdminLTE dashboard demo (This is only for demo purposes) -->
<script src="{{cdnjs "/static/adminlte//js/pages/dashboard3.js"}}"></script>
{{end}}

{{define "main_content_block"}}
<div class="row">
{{range $index,$item := .Lists}}
  <div class="col-md-2">
    {{if $item.IsPublic }}
    <div class="card card-primary card-outline">
    {{else}}
    <div class="card card-success card-outline">
    {{end}}
      
        <div class="card-header">
          <div class="row">
            <h3 class="card-title overflow-hidden text-truncate">
              {{if $item.IsPublic }}
              <button type="button" class="btn btn-primary btn-sm">{{i18n .Lang "common.public"}}</button>
              {{else}}
              <button type="button" class="btn btn-success btn-sm">{{i18n .Lang "common.private"}}</button>
              {{end}}
              {{$item.BookName}}
            </h3>
          </div>
          
          <div class="card-tools">
            <a href="{{urlfor "DocumentController.Edit" ":key" $item.Identify ":id" ""}}">
              <button type="button" class="btn btn-tool">
              <i class="fas fa-edit"></i>
              </button>
            </a>
            <button type="button" class="btn btn-tool" data-card-widget="card-refresh" data-source="#" data-source-selector="#card-refresh-content" data-load-on-init="false">
              <i class="fas fa-sync-alt"></i>
            </button>
            <button type="button" class="btn btn-tool" data-card-widget="collapse">
              <i class="fas fa-minus"></i>
            </button>
            
            <button type="button" class="btn btn-tool" data-card-widget="remove">
              <i class="fas fa-times"></i>
            </button>
          </div>
        </div>

      <div class="card-body">
        <a href="{{urlfor "DocumentController.Index" ":key" $item.Identify}}" title="{{$item.BookName}}-{{$item.CreateName}}">
        <img src="{{cdnimg $item.Cover}}" class="img-fluid" alt="{{$item.BookName}}-{{$item.CreateName}}" onerror="this.src='{{cdnimg "static/images/book.jpg"}}';">
        </a>
      </div>
    </div>
  </div>
{{end}}
</div>

{{end}}