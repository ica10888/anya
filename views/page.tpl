<!DOCTYPE html>

<html>
<head>
 <meta http-equiv="content-type" content="text/html; charset=utf-8" />
 <meta name="viewport" content="width=device-width, initial-scale=0.5" />

 <link rel="stylesheet" type="text/css" href="/static/css/github-markdown.css">
 <link rel="stylesheet" type="text/css" href="/static/css/main.css">

 <script type="text/javascript" src="https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js"></script>

 <link href="https://apps.bdimg.com/libs/highlight.js/9.1.0/styles/solarized_light.min.css" rel="stylesheet">
 <script src="http://apps.bdimg.com/libs/highlight.js/9.1.0/highlight.min.js"></script>

 <script src="/static/js/page.js"></script>
</head>
<body>

<title>{{ .Title }}</title>
<div style="height: 70px" class="inner">
</div>
<div>
 <article class="markdown-body">
  {{ .Content | str2html }}
 </article>
</div>
<div style="height: 70px" class="inner">
</div>
<div class="inner" id="doc-link">
</div>
<div class="inner" id="footer-mark">

</div>
<audio controls="controls" style="display: none;"></audio>
</body>
</html>