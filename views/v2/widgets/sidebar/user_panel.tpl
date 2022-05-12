<!-- Sidebar user panel (optional) -->
<div class="user-panel mt-3 pb-3 mb-3 d-flex">
    {{if gt .Member.MemberId 0}}
    <div class="image">
      <img src="{{cdnimg .Member.Avatar}}" class="img-circle elevation-2" alt="User Image">
    </div>
    <div class="info">
      <a href="" class="d-block">
        {{ if .Member.RealName }}
          {{.Member.RealName}}({{ .Member.Account }})
        {{else}}
          {{ .Member.Account }}
        {{end}}
      </a>
    </div>
    {{else}}
    <div class="image">
      <img src="/static/adminlte/img/user2-160x160.jpg" class="img-circle elevation-2" alt="User Image">
    </div>
    <div class="info">
      <a href="#" class="d-block">请登录</a>
    </div>
    {{end}}
</div>