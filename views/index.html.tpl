{{define "pagetitle"}} Index{{end}}

<div class="row">
	<img src="assets/fos2017.png" alt="fos2017 logo" class="img-responsive center-block" />
</div>
{{$loggedin := .loggedin}} {{if $loggedin}}
 <div class="row">
    <div class="col-md-offset-1 col-md-10">
        <h1 class="text-center" >Avant d'aller plus loin, vous devez avoir une photo prête à télécharger et connaître votre numéro de sécu.</h1>
    </div>
 </div>
 <div class="row">
    <div class="col-md-offset-1 col-md-10">
        <h2 class="text-center">Inscription réservée aux personnes majeures. Les personnes les plus disponibles seront contactées en priorité. Toute personne inscrite ne se sera pas automatiquement retenue.</h2>
    </div>
 </div>
 <div class="row">
    <div class="col-md-offset-5 col-md-2">
        <a href="/form" class="btn btn-info btn-block" role="button">J'ai compris</a>
    </div>
</div>
{{end}}