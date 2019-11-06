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

    <header id="top">
        <nav>
            <ul>
                <li><a href="/pages/about.html">
                        <img src="/static/assets/svg/mention.svg"><span>About</span></a></li>
                <li><a href="/">
                        <img src="/static/assets/svg/home.svg"><span>Home</span></a></li>
            </ul>
        </nav>
    </header>

    <div style="border-bottom: 1px solid #000"></div>


  <title>{{ .Title }}</title>
  <div class="inner">
    <h2>>>>> Title: {{ .Title }} <<<<</h2>
    <p>some description</p>
   </div>
   <div>
   <article class="markdown-body">
   {{ .Content | str2html }}
   </article>
   </div>
  <div class="inner" id="doc-link">
  </div>
  <audio controls="controls" style="display: none;"></audio>
  </body>
</html>