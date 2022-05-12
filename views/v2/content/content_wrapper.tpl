<div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <div class="content-header">
      <div class="container-fluid">
        <div class="row mb-2">
          <div class="col-sm-6">
            <h1 class="m-0">{{ .CONTENT_NAME }}</h1>
          </div><!-- /.col -->
          <div class="col-sm-6">
            <ol class="breadcrumb float-sm-right">
              <li class="breadcrumb-item"><a href="#">{{i18n .Lang "common.home" }}</a></li>
              <li class="breadcrumb-item active">{{ .CONTENT_NAME }}</li>
            </ol>
          </div><!-- /.col -->
        </div><!-- /.row -->
      </div><!-- /.container-fluid -->
    </div>
    <!-- /.content-header -->

    <!-- Main content -->
    <div class="content">
      <div class="container-fluid">
        {{ block "main_content_block" .}}
        
        {{ end }}
      </div>
      <!-- /.container-fluid -->
    </div>
    <!-- /.content -->
  </div>