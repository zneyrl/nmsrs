{{define "title"}}Register{{end}}

{{define "content"}}
<div class="container">
    <div class="row align-items-center">
        <div class="col-lg-3 mx-lg-auto">
            <a href="/template/home.tpl">
                <img src="/static/img/logo/dole-logo.png" height="192" class="mx-auto d-block logo" alt="Dole Logo">
            </a>
            <form>
                <div class="form-group">
                    <input type="text" class="form-control" name="name" placeholder="Name">
                </div>
                <div class="form-group">
                    <input type="text" class="form-control" name="username" placeholder="Username">
                </div>
                <div class="form-group">
                    <input type="password" class="form-control" name="password" placeholder="Password">
                </div>
                <div class="form-group">
                    <input type="password" class="form-control" name="confirm_password" placeholder="Confirm Password">
                </div>
                <button type="submit" class="btn btn-outline-primary btn-block">Register</button>
            </form>
        </div>
    </div>
</div>
{{end}}