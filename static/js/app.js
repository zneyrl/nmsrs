$(function () {
    $.fn.serializeObject = function () {
        var arr = this.serializeArray();
        var obj = {};

        for (var i = 0; i < arr.length; i++) {
            obj[arr[i].name] = arr[i].value;
        }
        return obj;
    }

    setCheckboxBoolValue = function (checkbox) {
        checkbox.on("change", function () {
            $(this).val($(this).is(":checked"));
        });
    }

    removeErrorMarkup = function (field) {
        var fc = field.parent();

        if (fc.hasClass("has-danger")) {
            fc.removeClass("has-danger");
        }

        if (field.hasClass("form-control-danger")) {
            field.removeClass("form-control-danger");
        }

        if (fc.find("div.form-control-feedback").get().length == 1) {
            fc.find("div.form-control-feedback").remove();
        }
    }

    addErrorMarkup = function (field, message) {
        var fc = field.parent();

        if (!fc.hasClass("has-danger")) {
            fc.addClass("has-danger");
        }

        if (!field.hasClass("form-control-danger")) {
            field.addClass("form-control-danger");
        }

        if (fc.find("div.form-control-feedback").get().length == 0) {
            fc.append(`<div class="form-control-feedback">` + message + `</div>`);
        }
    }

    quickRequest = function (action, method, data) {
        $.ajax({
            url: action,
            type: method,
            data: data,
            dataType: "json",
            error: function (r) {
                console.log(r);
            }
        }).done(function (r) {
            if (r.status == 200) {
                if (r.data.redirect != null) {
                    window.location.href = r.data.redirect;
                }
            }
            console.log(r)
        });
    }

    validateImage = function (elem, max_upload_size) {
        elem.on("change", function () {
            checkFileRequest("/check/file/image/" + $(this)[0].id, "POST", this)
            var preview = $(this).parent().find("#preview");
            var old_photo = preview.attr("src");

            if (this.files[0].size > max_upload_size) {
                preview.attr("src", old_photo);
                return
            } // TODO: First check file size and on server side

            if (this.files && this.files[0]) {
                var reader = new FileReader();

                reader.onload = function (e) {
                    preview.attr("src", e.target.result);

                    preview.on("error", function () {
                        preview.attr("src", old_photo);
                    });
                }
                reader.readAsDataURL(photo.files[0]);
            } else {
                preview.attr("src", old_photo);
            }
        });
    }

    checkFileRequest = function (url, method, input) {
        var alert = $("#alert");
        var data = new FormData();
        data.append(input.id, input.files[0]);

        $.ajax({
            url: url,
            type: method,
            data: data,
            dataType: "json",
            contentType: false,
            processData: false,
            success: function (r) {
                alert.empty();
                removeErrorMarkup($(input));
                errors = r.errors;

                try {
                    if (Object.keys(errors).length != 0) {
                        $.each(errors, function (k, v) {
                            var field = $("#" + k);
                            addErrorMarkup(field, v);
                        });
                    }
                } catch (e) {
                    var err_markup = `
                    <div class="alert alert-danger alert-dismissible fade show" role="alert">
                        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                        <i class="fa fa-exclamation-triangle"></i> `+ errors + `
                    </div>`;
                    alert.html(err_markup);
                }
            }, error: function (r) {
                console.log(r);
            }
        }).done(function (r) {
            console.log(r);
        });
    }

    makeRequest = function (form, fields, isMultipart) {
        var alert = $("#alert");
        var content_type = false;
        var data = (new FormData(form));
        var submitButton = $(form).find(":submit");
        var oldText = submitButton.text();
        submitButton.html(`<i class="fa fa-spinner fa-pulse fa-spin"></i> Please wait...`)

        if (!isMultipart) {
            content_type = "application/x-www-form-urlencoded; charset=UTF-8";
            data = $(form).serialize();
        }

        $.ajax({
            url: $(form).attr("action"),
            type: $(form).attr("method"),
            data: data,
            dataType: "json",
            contentType: content_type,
            processData: false,
            success: function (r) {
                alert.empty();

                $.each(fields, function (k, v) {
                    var field = $("#" + v);
                    removeErrorMarkup(field);
                });
                errors = r.errors;

                try {
                    if (Object.keys(errors).length != 0) {
                        $.each(errors, function (k, v) {
                            var field = $("#" + k);
                            addErrorMarkup(field, v);
                        });
                    }
                } catch (e) {
                    var err_markup = `
                    <div class="alert alert-danger alert-dismissible fade show" role="alert">
                        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                        <i class="fa fa-exclamation-triangle"></i> `+ errors + `
                    </div>`;
                    alert.html(err_markup);
                }
            }, error: function (r) {
                console.log(r);
            }
        }).done(function (r) {
            if (r.status == 200) {
                if (r.data.message != null) {
                    var message_markup = `
                    <div class="alert alert-success alert-dismissible fade show" role="alert">
                        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                        <i class="fa fa-check"></i> `+ r.data.message + `
                    </div>`;
                    alert.html(message_markup);
                }

                if (r.data.redirect != null) {
                    window.location.href = r.data.redirect;
                }

                $(form).find(":file").each(function () {
                    $(this).val("");
                });
            }
            submitButton.html(oldText);
            console.log(r);
        });
    }
});
