(function() {
    hljs.initHighlightingOnLoad();
    document.addEventListener('DOMContentLoaded', (event) => {
        document.querySelectorAll('.highlight pre').forEach((block) => {
            let odiv = block.parentElement.classList
            odiv.forEach((s) => {
                if(s.match("highlight-.*")){
                    let langs = s.replace("highlight-","")
                    hljs.configure({languages: [langs]});
                    hljs.highlightBlock(block);
                }})

        });
    });
})()


$.ajax({url:"/static/includes/head.html",success:function(result){
            $(result).prependTo($("body"))
}});

$.ajax({url:"/static/includes/footer.html",success:function(result){
        $("#footer-mark").after($(result))
    }});

$.ajax({url:"docList",success:function(result){

        let links = Object.keys(result)
        let titles = Object.values(result)
        let thisTitle = decodeURI(document.URL.split("/").pop())
        let index = links.lastIndexOf(thisTitle)
        if (index - 1 >= 0 ){
            $("<div>" +
                "<span class=\"prev-text\">上一篇： </span>" +
                "<a class=\"posts-list-name\" href=\""+ links[index - 1]+"\"> "+ titles[index - 1].title+"</a>" +
                "</div>"
            ).prependTo($("#doc-link"))
        }
        if (index + 1 < links.length ) {
            $("<div>" +
                "<span class=\"prev-text\">下一篇： </span>" +
                "<a class=\"posts-list-name\" href=\"" + links[index + 1] + "\"> " + titles[index + 1].title + "</a>" +
                "</div>"
            ).prependTo($("#doc-link"))
        }
}});


