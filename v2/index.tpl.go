package v2

var indexTpl = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>{{ .Title }} - Swagger UI</title>
  <link rel="icon" type="image/png" href="{{ .BasePath }}images/favicon-32x32.png" sizes="32x32" />
  <link rel="icon" type="image/png" href="{{ .BasePath }}images/favicon-16x16.png" sizes="16x16" />
  <link href='{{ .BasePath }}css/typography.css' media='screen' rel='stylesheet' type='text/css'/>
  <link href='{{ .BasePath }}css/reset.css' media='screen' rel='stylesheet' type='text/css'/>
  <link href='{{ .BasePath }}css/screen.css' media='screen' rel='stylesheet' type='text/css'/>
  <link href='{{ .BasePath }}css/reset.css' media='print' rel='stylesheet' type='text/css'/>
  <link href='{{ .BasePath }}css/print.css' media='print' rel='stylesheet' type='text/css'/>

  <script src='{{ .BasePath }}lib/object-assign-pollyfill.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/jquery-1.8.0.min.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/jquery.slideto.min.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/jquery.wiggle.min.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/jquery.ba-bbq.min.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/handlebars-2.0.0.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/js-yaml.min.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/lodash.min.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/backbone-min.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}swagger-ui.min.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/highlight.9.1.0.pack.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/highlight.9.1.0.pack_extended.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/jsoneditor.min.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/marked.js' type='text/javascript'></script>
  <script src='{{ .BasePath }}lib/swagger-oauth.js' type='text/javascript'></script>


  <!-- Some basic translations -->
  <!-- <script src='lang/translator.js' type='text/javascript'></script> -->
  <!-- <script src='lang/ru.js' type='text/javascript'></script> -->
  <!-- <script src='lang/en.js' type='text/javascript'></script> -->

  <script type="text/javascript">
    if (window.location.protocol == "https:") {
	window.location = "http://" + window.location.host + window.location.pathname;
    }
    $(function () {
     var cfg = {{ .ConfigJson }};
     var jsonEditorEnabled = false;
      {{if .JsonEditor }}
      if (window.localStorage && window.localStorage.getItem('jsonEditorEnabled')) {
        jsonEditorEnabled = true;
        $('#json-editor').text('Disable JSON Editor');
      } else {
        $('#json-editor').text('Enable JSON Editor');
      }
      {{end}}
      var url = window.location.protocol + "//" + window.location.host + "{{ .SwaggerJSON }}";

      hljs.configure({
        highlightSizeThreshold: 5000
      });

      // Pre load translate...
      if(window.SwaggerTranslator) {
        window.SwaggerTranslator.translate();
      }
      window.swaggerUi = new SwaggerUi({
        url: url,
        dom_id: "swagger-ui-container",
        onComplete: function(swaggerApi, swaggerUi){
          if(typeof initOAuth == "function") {
            initOAuth({
              clientId: "your-client-id",
              clientSecret: "your-client-secret-if-required",
              realm: "your-realms",
              appName: "your-app-name",
              scopeSeparator: ",",
              additionalQueryStringParams: {}
            });
          }

          if(window.SwaggerTranslator) {
            window.SwaggerTranslator.translate();
          }
        },
        onFailure: function(data) {
          log("Unable to Load SwaggerUI");
        },
        docExpansion: "none",
        jsonEditor: jsonEditorEnabled,
        apisSorter: "alpha",
        defaultModelRendering: 'schema',
        showRequestHeaders: false
      });

      function addApiKeyAuthorization(){
        return;
        var key = encodeURIComponent($('#input_apiKey')[0].value);
        if(key && key.trim() != "") {
            var apiKeyAuth = new SwaggerClient.ApiKeyAuthorization("api_key", key, "query");
            window.swaggerUi.api.clientAuthorizations.add("api_key", apiKeyAuth);
            log("added key " + key);
        }
      }

      $('#input_apiKey').change(addApiKeyAuthorization);

      // if you have an apiKey you would like to pre-populate on the page for demonstration purposes...
      /*
        var apiKey = "myApiKeyXXXX123456789";
        $('#input_apiKey').val(apiKey);
      */

      window.swaggerUi.load();

      function log() {
        if ('console' in window) {
          console.log.apply(console, arguments);
        }
      }
  });

    $(document).ready(function() {
      if (window.location.hash.length < 1) {
        setTimeout(function() {
          $(".toggleEndpointList").first().trigger("click");
        }, 500);
      }
    });

    function toggleJsonEditor() {
      if (window.localStorage) {
        if (window.localStorage.getItem('jsonEditorEnabled')) {
          window.localStorage.removeItem('jsonEditorEnabled');
        }
        else {
          window.localStorage.setItem('jsonEditorEnabled', "oh, yeah!");
        }
        window.location.reload();
      }
      else {
      	alert('Failed to access window.localStorage');
      }
    }
  </script>
</head>

<body class="swagger-section">
<div id='header'>
  <div class="swagger-ui-wrap">
    <a id="logo" href="">{{ .Title }}</a>
    <form id='api_selector'>
      <div class='input'><input placeholder="http://example.com/api" id="input_baseUrl" name="baseUrl" type="text"/></div>
      <div id='auth_container'></div>
      <div class='input'><a id="explore" class="header__btn" href="#" data-sw-translate>Explore</a></div>
    </form>
  </div>
</div>
<div class='input' style='margin:10px 0px 0 10px'>
{{if .JsonEditor }}<button id='json-editor' onclick="toggleJsonEditor()">Toggle JsonEditor</button>{{end}}
</div>


<div id="message-bar" class="swagger-ui-wrap" data-sw-translate>&nbsp;</div>
<div id="swagger-ui-container" class="swagger-ui-wrap"></div>
</body>
</html>
`
