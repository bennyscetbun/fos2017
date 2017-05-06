$(function () {
    function checkWhatYouWantToDo() {
        var newOutput = []
        $("#whatyouchoosetodo li").each(function (index, value) {
            newOutput.push($(value).val())
        })
        $("#WhatYouWantToDo").val(JSON.stringify(newOutput))
        $('#WhatYouWantToDo').trigger('input')
    }

    function initWhatYouWantToDo() {
        var valStr = $("#WhatYouWantToDo").val();
        var val = [];
        if (valStr.length > 2) {
            val = JSON.parse(valStr);
        }
        for (var i = 0; i < val.length; ++i) {
            $("#whatyoucando li[value='" + val[i] + "']").detach().appendTo('#whatyouchoosetodo');
        }
    }
    initWhatYouWantToDo();
    Sortable.create(whatyoucando, {
        group: "whatyouwantodo",
        // Element is dropped into the list from another list
        onAdd: function (/**Event*/evt) {
            checkWhatYouWantToDo()
        },

    });
    Sortable.create(whatyouchoosetodo, {
        group: "whatyouwantodo",
        // Element is dropped into the list from another list
        onAdd: function (/**Event*/evt) {
            checkWhatYouWantToDo()
        },
        onUpdate: function (/**Event*/evt) {
            checkWhatYouWantToDo()
        },
    });
});


$(function () {
    function initWhenCanYouBeThere() {
        var val = $("#WhenCanYouBeThere").val();
        if (val == "") {
            val = 0
        }
        $("#WhenCanYouBeThereChoices input").each(function (index, value) {
            if (($(value).val() & val) != 0) {
                $(value).prop('checked', true);
            }
        })
    }
    initWhenCanYouBeThere();

    $("#WhenCanYouBeThereChoices input").change(function () {
        var val = $("#WhenCanYouBeThere").val()
        if (val == "") {
            val = 0
        }
        if (this.checked) {
            val = val | $(this).val()
        } else {
            val = val & (~($(this).val()))
        }
        $("#WhenCanYouBeThere").val(val)
        $('#WhenCanYouBeThere').trigger('input')
    });
})


$(function () {
    $('form[data-toggle="validator"]').validator({
        custom: {
            whatyouwantodo: function ($el) {
                var valStr = $el.val()
                var val = []
                if (valStr.length > 2) {
                    val = JSON.parse(valStr)
                }
                val = val.filter(function (a) { return a !== 0; })
                if (val.length != 4) {
                    return "Vous devez mettre dans l ordre 4 choix de ce que vous voulez faire au Fall Of Summer"
                }
            },
            whencanyoubethere: function ($el) {
                var valStr = $el.val()
                if (valStr <= 0) {
                    return "Vous devez choisir au moins une date"
                }
            },
            checkdate: function ($el) {
                var valStr = $el.val()
                var parsedDate = moment(valStr, [
                    /*
                    "MM/DD/YYYY",
                    "DD/MM/YYYY",
                    "M/D/YYYY",
                    "D/M/YYYY",
                    "MM-DD-YYYY",
                    "M-D-YYYY",
                    "DD-MM-YYYY",
                    "D-M-YYYY",
                    */
                    "YYYY-MM-DD",
                    "YYYY-M-D",
                ], true)
                var diffYears = parsedDate.diff(moment(), 'years');
                if (!parsedDate.isValid() || diffYears < -100 || diffYears > -18) {
                    return "Veuillez écrire votre date au format année-mois-jour (ex: 1983-06-29)"
                }
            },
            checksecu: function ($el) {
                var pad = function(val, size) {
                    while (val.length < size) {
                        val = "0" + val
                    }
                    return val
                }

                var valStr = $el.val();
                if (valStr.length < 15) {
                    return "Veuillez entrer les 15 caractère qui composent votre numéro de sécurité social";
                }
                var re = new RegExp("^([1-37-8])([0-9]{2})(0[0-9]|[2-35-9][0-9]|[14][0-2])(?:(0[1-9]|[1-8][0-9]|9[0-69]|2[abAB])(00[1-9]|0[1-9][0-9]|[1-8][0-9]{2}|9[0-8][0-9]|990)|(9[78][0-9])(0[1-9]|[1-8][0-9]|90))(00[1-9]|0[1-9][0-9]|[1-9][0-9]{2})(0[1-9]|[1-8][0-9]|9[0-7])$");
                var result = valStr.match(re);
                if (result == null) {
                    return "votre numéro de sécurité social semble invalide";
                }
                var sexe = valStr.substring(0,1);
                var annee = pad(valStr.substring(1,3), 2);
                var mois = pad(valStr.substring(3,5), 2);
                var dept = pad(valStr.substring(5,7).replace("2A","19").replace("2B","18"), 2);
                var commune = pad(valStr.substring(7,10), 3);
                var ordre = pad(valStr.substring(10,13), 3);
                var cle = (valStr.substring(13,15));
                var numero=sexe+annee+mois+dept+commune+ordre;
                var calculCle = 97 - ( numero % 97);
                 if (calculCle != cle) {
                    return "votre numéro de sécurité social semble invalide";
                }
            }
        }
    })
})

$(function () {
    $("#photoupload input:button")
        .on('click', function () {
            $.ajax({
                // Your server script to process the upload
                url: '/upload',
                type: 'POST',

                // Form data
                data: new FormData($('form')[0]),

                // Tell jQuery not to process data or worry about content-type
                // You *must* include these options!
                cache: false,
                contentType: false,
                processData: false,

                // Custom XMLHttpRequest
                xhr: function () {
                    var myXhr = $.ajaxSettings.xhr();
                    return myXhr;
                },
                success: function (data) {
                    console.log(data)
                    d = new Date();
                    $('#photo').attr("src", data+"?"+d.getTime());
                    $('#photoHidden').val(data)
                    $('#photoHidden').trigger('input')

                },
                error: function () {
                    $('#notification-bar').text('An error occurred');
                }
            });
        });
});

$(function () {
    function getParameterByName(name, url) {
        if (!url) url = window.location.href;
        name = name.replace(/[\[\]]/g, "\\$&");
        var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
            results = regex.exec(url);
        if (!results) return null;
        if (!results[2]) return '';
        return decodeURIComponent(results[2].replace(/\+/g, " "));
    }
    var ok = getParameterByName('ok');

    if (ok != null) {
        $("#flash-success").text("Votre fiche a bien été enregistrée");
        $("#flash-success").show().delay(10000).fadeOut();

    }
});