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
    <div class="card card-success card-outline">
      
        <div class="card-header">
          <div class="row">
            <h3 class="col-10 card-title overflow-hidden text-truncate">
              <i class="fas fa-edit"></i>
              {{$item.BookName}}
            </h3>
            <div class="col-2 card-tools">
              <button type="button" class="btn btn-tool" data-card-widget="remove"><i class="fas fa-times"></i>
              </button>
            </div>
          </div>
        </div>

      <div class="card-body">
        <img src="{{cdnimg $item.Cover}}" class="img-fluid" alt="{{$item.BookName}}-{{$item.CreateName}}" onerror="this.src='{{cdnimg "static/images/book.jpg"}}';">
      </div>
    </div>
  </div>
{{end}}
</div>

{{end}}