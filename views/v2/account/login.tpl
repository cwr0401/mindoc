<!DOCTYPE html>
<html lang="zh-cn">
<head>
    
    {{template "v2/widgets/html_head.tpl" .}}

    <!-- icheck bootstrap -->
    <link rel="stylesheet" href="{{cdncss "/static/adminlte/plugins/icheck-bootstrap/icheck-bootstrap.min.css"}}">
    <!-- Theme style -->
    <link rel="stylesheet" href="{{cdncss "/static/adminlte/css/adminlte.min.css"}}">

</head>
<body class="hold-transition login-page">

<div class="login-box">
    <!-- /.login-logo -->
    <div class="card card-outline card-primary">
        <div class="card-header text-center">
            <a href="{{.BaseUrl}}" class="h1">{{.SITE_NAME}}</a>
        </div>
        <div class="card-body">
            <p class="login-box-msg">{{i18n .Lang "common.login"}}</p>

            <form role="form" method="post">
                {{ .xsrfdata }}
                <div class="input-group mb-3">
                    <input type="text" class="form-control" name="account" id="account" placeholder="{{i18n .Lang "common.email"}} / {{i18n .Lang "common.username"}}">
                    <div class="input-group-append">
                        <div class="input-group-text">
                            <span class="fas fa-envelope"></span>
                        </div>
                    </div>
                </div>
                <div class="input-group mb-3">
                    <input type="password" class="form-control" name="password" id="password" placeholder="{{i18n .Lang "common.password"}}">
                    <div class="input-group-append">
                        <div class="input-group-text">
                            <span class="fas fa-lock"></span>
                        </div>
                    </div>
                </div>

                {{if .ENABLED_CAPTCHA }}
                {{if ne .ENABLED_CAPTCHA "false"}}
                <div class="row mb-3">
                    <div class="col-6">
                        <input type="text" name="code" id="code" class="form-control"  maxlength="5" placeholder="{{i18n .Lang "common.captcha"}}" autocomplete="off">
                    </div>
                    <div class="col-6">
                        <img id="captcha-img" src="{{urlfor "AccountController.Captcha"}}" onclick="this.src='{{urlfor "AccountController.Captcha"}}?key=login&t='+(new Date()).getTime();" title={{i18n .Lang "message.click_to_change"}}>
                    </div>
                </div>
                {{end}}
                {{end}}

                <div class="row">
                    <div class="col-8">
                        <div class="icheck-primary">
                            <input type="checkbox" id="remember" name="is_remember">
                            <label for="remember">
                                {{i18n .Lang "common.keep_login"}}
                            </label>
                        </div>
                    </div>
                    <!-- /.col -->
                    <div class="col-4">
                        <button type="button" id="btn-login" class="btn btn-primary btn-block">Sign In</button>
                    </div>
                    <!-- /.col -->
                </div>
            </form>

            <p class="mb-1">
                <a href="forgot-password.html">{{i18n .Lang "common.forgot_password"}}</a>
            </p>
            {{if .ENABLED_REGISTER}}
            {{if ne .ENABLED_REGISTER "false"}}
            <p class="mb-0">
                <a href="register.html" class="text-center">{{i18n .Lang "common.register"}}</a>
            </p>
            {{end}}
            {{end}}
        </div>
        <!-- /.card-body -->
    </div>
    <!-- /.card -->

    <div id="alert-toast" role="alert" aria-live="assertive" aria-atomic="true" class="position-fixed toast bg-danger ml-1" data-delay="2000">
        <div class="toast-header">
          <i class="fas fa-bell mr-2 ml-2"></i>
          <strong class="mr-auto">Alert</strong>
          <button type="button" class="ml-2 mb-1 close" data-dismiss="toast" aria-label="Close">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>
        <div class="toast-body" id="toast-msg" style="min-width: 350px;">
        </div>
    </div>

</div>
<!-- /.login-box -->



<!-- jQuery -->
<script src="{{cdnjs "/static/adminlte/plugins/jquery/jquery.min.js"}}" type="text/javascript"></script>
<!-- Bootstrap 4 -->
<script src="{{cdnjs "/static/adminlte/plugins/bootstrap/js/bootstrap.bundle.min.js"}}" type="text/javascript"></script>
<!-- AdminLTE App -->
<script src="{{cdnjs "/static/adminlte/js/adminlte.min.js"}}" type="text/javascript"></script>

<script type="text/javascript">
    $(function () {
        $("#account,#password,#code").on('focus', function () {
            $(this).tooltip('dispose').parents('.form-group').removeClass('has-error');
        });

        $(document).keydown(function (e) {
            var event = document.all ? window.event : e;
            if (event.keyCode === 13) {
                $("#btn-login").click();
            }
        });

        $("#btn-login").on('click', function () {
            $(this).tooltip('dispose').parents('.form-group').removeClass('has-error');
            var $btn = $(this).button('loading');

            var account = $.trim($("#account").val());
            var password = $.trim($("#password").val());
            var code = $("#code").val();

            if (account === "") {
                $("#account").tooltip({ placement: "auto", title: "{{i18n .Lang "message.account_empty"}}", trigger: 'manual' })
                    .tooltip('show')
                    .parents('.form-group').addClass('has-error');
                $btn.button('reset');
                return false;
            } else if (password === "") {
                $("#password").tooltip({ title: '{{i18n .Lang "message.password_empty"}}', trigger: 'manual' })
                    .tooltip('show')
                    .parents('.form-group').addClass('has-error');
                $btn.button('reset');
                return false;
            } else if (code !== undefined && code === "") {
                $("#code").tooltip({ title: '{{i18n .Lang "message.captcha_empty"}}', trigger: 'manual' })
                    .tooltip('show')
                    .parents('.form-group').addClass('has-error');
                $btn.button('reset');
                return false;
            } else {
                $.ajax({
                    url: "{{urlfor "AccountController.Login" "url" .url}}",
                    data: $("form").serializeArray(),
                    dataType: "json",
                    type: "POST",
                    success: function (res) {
                        if (res.errcode !== 0) {
                            $("#captcha-img").click();
                            $("#code").val('');
                            // layer.msg(res.message);
                            $('#toast-msg').text(res.message)
                            $('#alert-toast').toast('show')
                            $btn.button('reset');
                        } else {
                            turl = res.data;
                            if (turl === "") {
                                turl = "/";
                            }
                            window.location = turl;
                        }
                    },
                    error: function () {
                        $("#captcha-img").click();
                        $("#code").val('');
                        // layer.msg('{{i18n .Lang "message.system_error"}}');
                        $('#toast-msg').text('{{i18n .Lang "message.system_error"}}')
                        $('#alert-toast').toast('show')
                        $btn.button('reset');
                    }
                });
            }

            return false;
        });
    });
</script>

</body>

</html>