<!DOCTYPE html>
<html lang="en">
<head>
    
{{template "v2/widgets/html_head.tpl" .}}

  <!-- IonIcons -->
  <link rel="stylesheet" href="https://code.ionicframework.com/ionicons/2.0.1/css/ionicons.min.css">
  <!-- Theme style -->
  <link rel="stylesheet" href="{{cdncss "/static/adminlte/css/adminlte.min.css"}}">

  <!-- OPTIONAL CCS -->
{{block "optional_css_block" .}}
{{end}}

</head>
<!--
`body` tag options:

  Apply one or more of the following classes to to the body tag
  to get the desired effect

  * sidebar-collapse
  * sidebar-mini
-->
<body class="hold-transition sidebar-mini">
<div class="wrapper">
  <!-- Navbar -->
{{template "v2/widgets/header.tpl" .}}
  <!-- /.navbar -->

  <!-- Main Sidebar Container -->
  <aside class="main-sidebar sidebar-dark-primary elevation-4">
    <!-- Brand Logo -->
    <a href="#" class="brand-link">
      <img src="{{cdnimg "/favicon.ico"}}" alt="Logo" class="brand-image img-circle elevation-3" style="opacity: .8">
      <span class="brand-text font-weight-light">{{.SITE_NAME}}</span>
    </a>

    <!-- Sidebar -->
    <div class="sidebar">
      <!-- Sidebar user panel (optional) -->
{{template "v2/widgets/sidebar/user_panel.tpl" .}}

      <!-- SidebarSearch Form -->
      <div class="form-inline">
        <div class="input-group" data-widget="sidebar-search">
          <input class="form-control form-control-sidebar" type="search" placeholder="Search" aria-label="Search">
          <div class="input-group-append">
            <button class="btn btn-sidebar">
              <i class="fas fa-search fa-fw"></i>
            </button>
          </div>
        </div>
      </div>

      <!-- Sidebar Menu -->
{{template "v2/widgets/sidebar/menu.tpl" .}}
      <!-- /.sidebar-menu -->
    </div>
    <!-- /.sidebar -->
  </aside>

  <!-- Content Wrapper. Contains page content -->
{{template "v2/content/content_wrapper.tpl" .}}
  <!-- /.content-wrapper -->

  <!-- Control Sidebar -->
  <aside class="control-sidebar control-sidebar-dark">
    <!-- Control sidebar content goes here -->
  </aside>
  <!-- /.control-sidebar -->

  <!--footer-->
{{template "v2/widgets/footer.tpl"}}

</div>
<!-- ./wrapper -->

<!-- REQUIRED SCRIPTS -->

<!-- jQuery -->
<script src="{{cdnjs "/static/adminlte/plugins/jquery/jquery.min.js"}}"></script>
<!-- Bootstrap -->
<script src="{{cdnjs "/static/adminlte/plugins/bootstrap/js/bootstrap.bundle.min.js"}}"></script>
<!-- AdminLTE -->
<script src="{{cdnjs "/static/adminlte/js/adminlte.js"}}"></script>

<!-- OPTIONAL SCRIPTS -->
{{block "optional_scripts_block" .}}
{{end}}
</body>
</html>
