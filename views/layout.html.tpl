<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">

	<title>{{template "pagetitle" .}}</title>

	<link rel="stylesheet" href="/assets/bootstrap.min.css" />
	<link rel="stylesheet" href="//maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css" />
	<link rel="stylesheet" href="/assets/fos2017.css" />
	<link rel="stylesheet" href="https://cdn.datatables.net/1.10.15/css/dataTables.bootstrap.min.css" />

	<script type="text/javascript" src="https://code.jquery.com/jquery-2.1.3.min.js"></script>
	<script type="text/javascript" src="//maxcdn.bootstrapcdn.com/bootstrap/3.3.2/js/bootstrap.min.js"></script>
	<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/1000hz-bootstrap-validator/0.11.9/validator.min.js"></script>
	<script type="text/javascript" src="https://cdn.datatables.net/1.10.15/js/jquery.dataTables.min.js"></script>
	<script type="text/javascript" src="https://cdn.datatables.net/1.10.15/js/dataTables.bootstrap.min.js"></script>
	<script type="text/javascript" src="/assets/moment.js"></script>
	<script type="text/javascript" src="/assets/sortable.js"></script>
</head>
<body class="container-fluid" style="padding-top: 15px;">
	<nav class="navbar navbar-default">
		<div class="container-fluid">
			<div class="navbar-header">
				<a class="navbar-brand" href="/">Fall Of Summer 2017: Portail bénévole</a>
			</div>

			<div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
				<ul class="nav navbar-nav navbar-right">
					{{if not .loggedin}}
					<li><a href="/auth/register">Devenir bénévole</a></li>
					<li><a href="/auth/login"><i class="fa fa-sign-in"></i> Déjà inscrit</a></li>
					{{else}}
					<li class="dropdown">
						<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">Welcome {{.current_user_name}}! <span class="caret"></span></a>
						<ul class="dropdown-menu" role="menu">
							<li>
								<a href="/auth/logout">
									<i class="fa fa-sign-out"></i> Logout
								</a>
							</li>
						</ul>
					</li>
					{{end}}
				</ul>
			</div>
		</div>
	</nav>

	<div id="flash-success" class="alert alert-success" style="display: none;"></div>
	{{with .flash_success}}<div class="alert alert-success">{{.}}</div>{{end}}
	{{with .flash_error}}<div class="alert alert-danger">{{.}}</div>{{end}}
	{{template "yield" .}}
	{{template "authboss" .}}
</body>
</html>
{{define "pagetitle"}}{{end}}
{{define "yield"}}{{end}}
{{define "authboss"}}{{end}}