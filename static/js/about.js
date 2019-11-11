$.ajax({url:"/static/includes/head.html",success:function(result){
        $(result).prependTo($("body"))
    }});

$.ajax({url:"/static/includes/footer.html",success:function(result){
        $("#footer-mark").after($(result))
    }});
