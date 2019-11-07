$.ajax({url:"/static/includes/head.html",success:function(result){
        $(result).prependTo($("body"))
    }});

$.ajax({url:"/static/includes/footer.html",success:function(result){
        $("#footer-mark").after($(result))
    }});


$.ajax({url:"/doc/docList",success:function(result){

        let links = Object.keys(result)
        let titles = Object.values(result)

        titles.forEach((title,index,arr) =>{
            $(
                "  <article class=\"doc-list-class\">\n" +
                "    <section class=\"inner\" id=\"doc-list\">" +
                "<header class=\"post-header\">\n" +
                "    <h1 class=\"post-title\"><a class=\"post-link\" href=\"doc/"+ links[index]+"\">"+ titles[index].title+"</a></h1>\n" +
                "    <div class=\"post-meta\">\n" +
                "      <span class=\"post-time\"> Posted by "+ titles[index].author +" on "+ titles[index].date +" </span>\n" +
                "      <div class=\"post-category\">\n" +
                "          </div>\n" +
                "    </div>\n" +
                "  </header>" +
                "    </section>\n" +
                "  </article>"

            ).prependTo($("#doc-list"))
        })
    }});