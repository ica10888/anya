<html>
 <head>
  <meta http-equiv="content-type" content="text/html; charset=utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=0.5" />
  <link rel="stylesheet" type="text/css" href="/static/css/github-markdown.css">
 </head>
 <body>
  <title>{{.Tittle}}</title>
  <div class="inner">
    <h2>>>>> Tittle: {{.Tittle}} <<<<</h2>
    <p>some description</p>
   </div>
   <div>
   <article class="markdown-body">
   {{ .Content | str2html}}
   </article>
   </div>
  <audio controls="controls" style="display: none;"></audio>
  </body>
</html>