{{template "v2/base.tpl" .}}

{{define "optional_css_block"}}
  <!-- Tempusdominus Bootstrap 4 -->
  <link rel="stylesheet" href="{{cdncss "/static/adminlte/plugins/tempusdominus-bootstrap-4/css/tempusdominus-bootstrap-4.min.css"}}">

  <!-- Daterange picker -->
  <link rel="stylesheet" href="{{cdncss "/static/adminlte/plugins/daterangepicker/daterangepicker.css"}}">
{{end}}

{{define "optional_scripts_block"}}
<script src="{{cdnjs "/static/adminlte/plugins/jquery-ui/jquery-ui.min.js"}}"></script>
<script>
  $.widget.bridge('uibutton', $.ui.button)
</script>
<!-- AdminLTE for demo purposes -->
<script src="{{cdnjs "/static/adminlte/js/demo.js"}}"></script>
<script src="{{cdnjs "/static/jquery-plugins/datetime/datetime.js"}}"></script>

<!-- daterangepicker -->
<script src="{{cdnjs "/static/adminlte/plugins/moment/moment.min.js"}}"></script>
<script src="{{cdnjs "/static/adminlte/plugins/daterangepicker/daterangepicker.js"}}"></script>
<!-- Tempusdominus Bootstrap 4 -->
<script src="{{cdnjs "/static/adminlte/plugins/tempusdominus-bootstrap-4/js/tempusdominus-bootstrap-4.min.js"}}"></script>

<script>
    
    function change_datetime() {
      var now = new Date();
      //date格式化方法
      var nowStr = DateUtil.dateToStr("yyyy-MM-dd HH:mm:ss", now);
      var week = DateUtil.getWeek(now);
      $('#datetime').text(nowStr + " 星期" + week);
    }

    function change_wkzone(wkinfo){
      if (wkinfo.errcode !== 200) {
        return
      }
      $('#appPrimaryZone').empty()
      wkinfo.data.app_primary.forEach(element => {

        card = render_wkzone_card(element, "primary");
        $('#appPrimaryZone').append(card);

      });
      $('#platformZone').empty();
      carouselItems = ''
      wkinfo.data.platform.forEach((element, index, arr) => {
        card = render_wkzone_card(element, "success");
        if (index % 2 == 0) {
          if (index === 0) {
            carouselItems = carouselItems + '<div class="carousel-item active" data-interval="5000">' + card;
          } else {
            carouselItems = carouselItems + '<div class="carousel-item" data-interval="5000">' + card; 
          }
        } else {
          carouselItems = carouselItems + card + '</div>';
        }
      });
      if (wkinfo.data.platform.length % 2 == 1) {
        carouselItems += '</div>';
      }
      carousel = create_carousel_interval('carouselPlatformInterval1', carouselItems);
      $('#platformZone').append(carousel);
      $('#carouselPlatformInterval1').carousel();
      $('#teamLeaderZone').empty()
      wkinfo.data.team_leader.forEach(element => {
        card = render_wkzone_card(element, "primary");
        $('#teamLeaderZone').append(card);
      });
    }

    function render_wkzone_card(wk, level){
      labels = '';
      wk.labels.forEach( label => {
        labels += '<span class="badge badge-' + level + ' float-right ml-1">' + label + '</span>'
      });
      cardHeader = '<div class="card-header">' +
        '<h3 class="card-title">' +
        '<icon class="fa fa-bell mr-2"></icon>' +
        wk.title +
        '</h3>' +
        labels +
        '</div>';
      cardBody = '<div class="card-body box-profile">' +
        '<h3 class="profile-username text-center">' + wk.staff_name + '</h3>' +
        '<p class="text-muted text-center">' + wk.staff_phone + '<br/>'  +  wk.staff_email + '</p>' +
        '<ul class="list-group list-group-unbordered mb-3">' +
        '<li class="list-group-item">' +
        '<b>地点</b> <a class="float-right">'+ wk.position + '</a>' +
        '</li>' +
        '<li class="list-group-item">' +
        '<b>时段</b> <a class="float-right">' + wk.start_time.substr(5,5) + ' ' + wk.start_time.substr(11,5) + ' ~ ' + wk.end_time.substr(5, 5) + ' ' + wk.end_time.substr(11, 5) + '</a>' +
        '</li>' +
        '</ul>' +
        '</div>';

      card = '<div class="card card-' + level + ' card-outline">' +
        cardHeader +
        cardBody +
        '</div>';
      return card;
    }

    function create_carousel_interval(id, items){
      carousel = '<div id="' + id + '" class="carousel slide" data-ride="carousel">' +
        '<div class="carousel-inner">' +
          items +
        '</div>' +
        '<button class="carousel-control-prev" type="button" data-target="#carouselPlatformInterval" data-slide="prev">' +
          '<span class="carousel-control-prev-icon" aria-hidden="true"></span>' +
          '<span class="sr-only">Previous</span>' +
        '</button>' +
        '<button class="carousel-control-next" type="button" data-target="#carouselPlatformInterval" data-slide="next">' +
          '<span class="carousel-control-next-icon" aria-hidden="true"></span>' +
          '<span class="sr-only">Next</span>' +
        '</button>' +
      '</div>';
      return carousel;
    }

    function watchkeeper() {
      $.get(
        "{{urlfor "WKController.GetWatchKeeperInfo"}}",
        function (data, status) {
          oldWkInfo = sessionStorage.getItem("wkinfo");
          newWkInfo = JSON.stringify(data);
          // localStorage.setItem("wkinfo", JSON.stringify(data))
          // console.log(JSON.stringify(data))
          if (oldWkInfo !== newWkInfo) {
            sessionStorage.setItem("wkinfo", newWkInfo);
            change_wkzone(data);
            console.log("update watchkeeper dashboard")
          }
          
          console.log("get watchkeeper info");
        }
      )
    }

    $(function(){
        $('#carouselPlatformInterval').carousel();
        change_datetime();
        MyClock = setInterval(change_datetime, 1000);

        // $('#calendar').datetimepicker({
        //     locale: 'zh-cn',
        //     format: 'L',
        //     inline: true
        // })
        
        sessionStorage.removeItem("wkinfo")
        watchkeeper();
        setInterval(watchkeeper, 60 * 1000);
        
    })
    
    
</script>


{{end}}

{{define "main_content_block"}}
<h4>
  <span class="ml-2" id="datetime">日期时间</span>
  <span class="ml-5">值班电话：</span><a class="ml-2">15927428738</a>
  <span class="ml-3">团队邮箱：</span><a class="ml-2">tsc.zb@ccbft.com</a>
</h4>
<div class="row">
    <div id="appPrimaryZone" class="col-md-2">
      {{range $index, $item := .app_primary}}
      <div class="card card-primary card-outline">
        <div class="card-header">     
            <h3 class="card-title">
              <icon class="fa fa-bell mr-2"></icon>
                {{ $item.Title }}
            </h3>
            {{range $lindex, $label := $item.Labels }}
              <span class="badge badge-primary float-right ml-1">{{ $label }}</span>
            {{end}}
          </div>
          <div class="card-body box-profile">
                <h3 class="profile-username text-center">{{ $item.StaffName }}</h3>

                <p class="text-muted text-center">{{ $item.StaffPhone }} <br/> {{ $item.StaffEmail }}</p>

                <ul class="list-group list-group-unbordered mb-3">
                  <li class="list-group-item">
                    <b>地点</b> <a class="float-right">{{ $item.Position}}</a>
                  </li>
                  <li class="list-group-item">
                    <b>时段</b> <a class="float-right">{{ dateformat $item.StartTime "01-02 15:04"}} ~ {{ dateformat $item.EndTime "01-02 15:04" }}</a>
                  </li>
                </ul>
          </div>
        </div>
      {{else}}
          <h3>空空如也</h3>
      {{end}}
    </div>
    <div id="platformZone" class="col-md-2">
      
      <div id="carouselPlatformInterval" class="carousel slide" data-ride="carousel">
        <div class="carousel-inner">
        {{range $index, $elem := .platform}}
          {{ if eq (mod $index 2) 0}} 
            {{ if eq $index 0 }}
          <!-- carousel-item -->
          <div class="carousel-item active" data-interval="5000">
            {{else}}
          <div class="carousel-item" data-interval="5000">
            {{end}}
          {{end}}
            <div class="card card-success card-outline">
              <div class="card-header">
                <h3 class="card-title">
                    <icon class="fa fa-bell mr-2"></icon>
                    {{ $elem.Title }}
                </h3>
                {{range $lindex, $label := $elem.Labels }}
                <span class="badge badge-success float-right ml-1">{{ $label }}</span>
                {{end}}
              </div>
              <div class="card-body">
              <h3 class="profile-username text-center">{{ $elem.StaffName }}</h3>
              <p class="text-muted text-center">{{ $elem.StaffPhone }} <br/> {{ $elem.StaffEmail }}</p>
              <ul class="list-group list-group-unbordered mb-3">
                <li class="list-group-item">
                  <b>地点</b> <a class="float-right">{{ $elem.Position }}</a>
                </li>
                <li class="list-group-item">
                  <b>时段</b> <a class="float-right">08:30 ~ 20:30</a>  
                </li>
              </ul>
              </div>
            </div>
          {{ if eq (mod $index 2) 1}}
          </div>
          <!-- ./carousel-item -->
          {{end}}
        {{end}}
        {{ $elem_len_mod := mod (len .platform) 2}}
        {{ if eq $elem_len_mod 1 }}
          </div>
          <!-- carousel-item -->
        {{ end }}

        </div>
        <button class="carousel-control-prev" type="button" data-target="#carouselPlatformInterval" data-slide="prev">
          <span class="carousel-control-prev-icon" aria-hidden="true"></span>
          <span class="sr-only">Previous</span>
        </button>
        <button class="carousel-control-next" type="button" data-target="#carouselPlatformInterval" data-slide="next">
          <span class="carousel-control-next-icon" aria-hidden="true"></span>
          <span class="sr-only">Next</span>
        </button>
      </div>

    </div>
    <div id="teamLeaderZone" class="col-md-2">
      {{range $index,$item := .team_leader}}
        <div class="card card-primary card-outline">
        <div class="card-header">
            <h3 class="card-title">
                <icon class="fa fa-bell mr-2"></icon>
                {{ $item.Title }}
            </h3>
        </div>
        <div class="card-body box-profile">
            <h3 class="profile-username text-center">{{ $item.StaffName }}</h3>

            <p class="text-muted text-center">{{ $item.StaffPhone }} <br/> {{ $item.StaffEmail }}</p>

        </div>
        </div>
      {{end}}
        <div class="card card-primary card-outline">
            <div class="card-header">
                <h3 class="card-title">
                    <icon class="fa fa-bell mr-2"></icon>
                    一线组长
                </h3>    
            </div>
            <div class="card-body">
                <h3 class="profile-username text-center">万仁亮</h3>
  
                <p class="text-muted text-center">18164003738 <br/> wanrenliang.zb@ccbft.com</p>
            </div>
        </div>
    </div>
    <div class="col-md-6" id="testZone">


    </div>
</div>

{{end}}