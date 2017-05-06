 {{define "pagetitle"}} form{{end}}

{{$loggedin := .loggedin}} {{if $loggedin}}
<script type="text/javascript" src="/assets/fos2017.js"></script>
<div class="row">


	<div class="col-md-offset-1 col-md-10">
		<form id="photoupload" enctype="multipart/form-data" action="/upload" method="post">
			<div class="form-group has-feedback">
				<label for="Photo" class="control-label">Photo d'identité*</label>
				<input type="file" name="uploadfile" />
				<input type="hidden" name="csrf_token" value="{{.csrf_token}}" />
				<input type="button" class="btn btn-warning" value="Après avoir choisi le fichier de la photo, Cliquez ici pour envoyer votre photo. Si votre photo ne s'affiche pas en dessous de ce bouton, vous ne pourrez pas valider le formulaire" />
			</div>
		</form>
		<img id="photo" {{if .picture }} src={{.picture}} {{end}} alt="No Photo" style="max-width:300px;max-height:300px;">
	</div>
</div>
<div class="row">
	<div class="col-md-offset-1 col-md-10">
		<form data-toggle="validator" method="POST" action="/form" role="form">
			<div class="form-group has-feedback">
				<input type="hidden" class="form-control" id="photoHidden" name="Photo"
				data-validate="true"
				data-required-error="Veuillez envoyer une photo de vous"
				 {{if .userinfo.Photo}} value={{.userinfo.Photo}} {{end}}
				required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="Firstname" class="control-label">Prénom*</label>
				<input type="text" class="form-control" id="Firstname" name="Firstname" placeholder="Prénom" data-required-error="Veuillez remplir ce champ"
				 {{if .userinfo.Firstname}} value={{.userinfo.Firstname}} {{end}} required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="Lastname" class="control-label">Nom*</label>
				<input type="text" class="form-control" id="Lastname" name="Lastname" placeholder="Nom" data-required-error="Veuillez remplir ce champ"
				 {{if .userinfo.Lastname}} value={{.userinfo.Lastname}} {{end}} required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="Address" class="control-label">Adresse*</label>
				<input type="text" class="form-control" id="Address" name="Address" placeholder="Adresse" data-required-error="Veuillez remplir ce champ"
				 {{if .userinfo.Address}} value={{.userinfo.Address}} {{end}} required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="CP" class="control-label">Code Postal*</label>
				<input type="text" class="form-control" id="CP" name="CP" placeholder="Code Postal" data-required-error="Veuillez remplir ce champ"
				 {{if .userinfo.CP}} value={{.userinfo.CP}} {{end}} required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="Town" class="control-label">Ville*</label>
				<input type="text" class="form-control" id="Town" name="Town" placeholder="Ville" data-required-error="Veuillez remplir ce champ"
				 {{if .userinfo.Town}} value={{.userinfo.Town}} {{end}} required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="PhoneNumber" class="control-label">Numéro de portable*</label>
				<input type="text" pattern="^[+]?[0-9]{1,}$" class="form-control" id="PhoneNumber" name="PhoneNumber" placeholder="Numéro de portable"
				 data-pattern-error="Ce numéro semble incorrect" data-required-error="Veuillez remplir ce champ" {{if .userinfo.PhoneNumber}}
				 value={{.userinfo.PhoneNumber}} {{end}} required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="HealthNumber" class="control-label">Numéro de sécu*</label>
				<input type="text"  class="form-control" id="HealthNumber" name="HealthNumber" placeholder="Numéro de sécurité sociale"
				 maxlength="15" data-pattern-error="Ce numéro de sécurité sociale ne semble pas valide" data-required-error="Veuillez remplir ce champ"
				 {{if .userinfo.HealthNumber}} value={{.userinfo.HealthNumber}} {{end}}
				 data-checksecu
				 required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="BirthDate" class="control-label">Date de naissance*</label>
				<input type="text" class="form-control" id="BirthDate" name="BirthDate" placeholder="année-mois-jour" data-required-error="Veuillez remplir ce champ"
				 {{if gt .birthdate "1910-01-01"}} value={{.birthdate}} {{end}}
				 data-checkdate
				 required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="BirthPlace" class="control-label">Lieu de naissance*</label>
				<input type="text" class="form-control" id="BirthPlace" name="BirthPlace" placeholder="Lieu de naissance" data-required-error="Veuillez remplir ce champ"
				 {{if .userinfo.BirthPlace}} value={{.userinfo.BirthPlace}} {{end}} required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>

			<div class="form-group has-feedback">
				<label for="EmergencyContact" class="control-label">Contact en cas d'urgence*</label>
				<div class="form-group has-feedback">
					<select class="form-control" id="EmergencyContactType" name="EmergencyContactType" data-required-error="Veuillez indiquer le lien que vous avez avec votre contact en cas d urgence" required>
						<option value="" disabled {{ if eq .userinfo.EmergencyContactType 0}} selected {{end}}> Lien avec vous</option>
						<option value="1" {{ if eq .userinfo.EmergencyContactType 1}} selected {{end}}>Père</option>
						<option value="2" {{ if eq .userinfo.EmergencyContactType 2}} selected {{end}}>Mère</option>
						<option value="3" {{ if eq .userinfo.EmergencyContactType 3}} selected {{end}}>Famille</option>
						<option value="4" {{ if eq .userinfo.EmergencyContactType 4}} selected {{end}}>Compagnon</option>
						<option value="5" {{ if eq .userinfo.EmergencyContactType 5}} selected {{end}}>Autre</option>
					</select>
					<div class="help-block with-errors"></div>
				</div>
				<div class="form-group has-feedback">
					<input type="text" class="form-control" id="EmergencyContactFirstname" name="EmergencyContactFirstname" placeholder="Prénom*" data-required-error="Veuillez remplir ce champ"
						{{if .userinfo.EmergencyContactFirstname}} value={{.userinfo.EmergencyContactFirstname}} {{end}} required>
					<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
					<div class="help-block with-errors"></div>
				</div>
				<div class="form-group has-feedback">
					<input type="text" class="form-control" id="EmergencyContactLastname" name="EmergencyContactLastname" placeholder="Nom*" data-required-error="Veuillez remplir ce champ"
						{{if .userinfo.EmergencyContactLastname}} value={{.userinfo.EmergencyContactLastname}} {{end}} required>
					<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
					<div class="help-block with-errors"></div>
				</div>
				<div class="form-group has-feedback">
					<input type="text" class="form-control" pattern="^[+]?[0-9]{1,}$" id="EmergencyContactPhoneNumber" name="EmergencyContactPhoneNumber" placeholder="Numéro de téléphone*" data-required-error="Veuillez remplir ce champ"
						{{if .userinfo.EmergencyContactPhoneNumber}} value={{.userinfo.EmergencyContactPhoneNumber}} {{end}}
						data-pattern-error="Ce numéro semble incorrect"
						required>
					<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
					<div class="help-block with-errors"></div>
				</div>
				<div class="form-group has-feedback">
					<input type="text" class="form-control" id="EmergencyContactAddress" name="EmergencyContactAddress" placeholder="Adresse" data-required-error="Veuillez remplir ce champ"
						{{if .userinfo.EmergencyContactAddress}} value={{.userinfo.EmergencyContactAddress}} {{end}}>
					<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
					<div class="help-block with-errors"></div>
				</div>
				<div class="form-group has-feedback">
					<input type="text" class="form-control" id="EmergencyContactCP" name="EmergencyContactCP" placeholder="Code Postal" data-required-error="Veuillez remplir ce champ"
						{{if .userinfo.EmergencyContactCP}} value={{.userinfo.EmergencyContactCP}} {{end}}>
					<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
					<div class="help-block with-errors"></div>
				</div>
				<div class="form-group has-feedback">
					<input type="text" class="form-control" id="EmergencyContactTown" name="EmergencyContactTown" placeholder="Ville" data-required-error="Veuillez remplir ce champ"
						{{if .userinfo.EmergencyContactTown}} value={{.userinfo.EmergencyContactTown}} {{end}}>
					<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
					<div class="help-block with-errors"></div>
				</div>
			</div>
			<div class="form-group has-feedback">
				<label for="Facebook" class="control-label">Pseudo Facebook</label>
				<input type="text" class="form-control" id="Facebook" name="Facebook" placeholder="Pseudo Facebook pour vous rajouter au group"
				 {{if .userinfo.Facebook}} value={{.userinfo.Facebook}} {{end}}>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="TShirt" class="control-label">Taille de Tshirt*</label>
				<div> </div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="1"
						{{if eq .userinfo.TShirt 1}}
							checked
						{{end}}
					required>Homme S</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="2"
					{{if eq .userinfo.TShirt 2}}
							checked
					{{end}}
					required>Homme M</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="3"
					{{if eq .userinfo.TShirt 3}}
							checked
					{{end}}
					required>Homme L</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="4"
					{{if eq .userinfo.TShirt 4}}
							checked
					{{end}}
					required>Homme XL</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="5"
					{{if eq .userinfo.TShirt 5}}
							checked
					{{end}}
					required>Homme XXL</label>
				</div>
				<div> </div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="6"
					{{if eq .userinfo.TShirt 6}}
						checked
					{{end}}
					required>Girly S</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="7"
					{{if eq .userinfo.TShirt 7}}
							checked
					{{end}}
					required>Girly M</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="8"
					{{if eq .userinfo.TShirt 8}}
							checked
					{{end}}
					required>Girly L</label>
				</div>
				<div> </div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="9"
					{{if eq .userinfo.TShirt 9}}
							checked
					{{end}}
					required>Tank Top S</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="10"
					{{if eq .userinfo.TShirt 10}}
							checked
					{{end}}
					required>Tank Top M</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="11"
					{{if eq .userinfo.TShirt 11}}
							checked
					{{end}}
					required>Tank Top L</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="TShirt" value="12"
					{{if eq .userinfo.TShirt 12}}
							checked
					{{end}}
					required>Tank Top XL</label>
				</div>
				<div class="help-block with-errors"></div>
			</div>

			<div class="form-group has-feedback">
				<label for="Regime" class="control-label">Régime alimentaire*</label>
				<select class="form-control" id="Regime" name="Regime" data-required-error="Veuillez Choisir un régime alimentaire" required>
					<option value="" disabled {{ if eq .userinfo.Regime 0}} selected {{end}}> Régime alimentaire</option>
					<option value="1" {{ if eq .userinfo.Regime 1}} selected {{end}}>Omni(sans régime spécifique)</option>
					<option value="2" {{ if eq .userinfo.Regime 2}} selected {{end}}>Végétarien</option>
					<option value="3" {{ if eq .userinfo.Regime 3}} selected {{end}}>Végétalien</option>
				</select>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="Allergy" class="control-label">Avez-vous des allergies graves. Si oui, lesquelles ?</label>
				<input type="text" class="form-control" id="Allergy" name="Allergy" placeholder="Allergie" data-required-error="Veuillez remplir ce champ"
				 {{if .userinfo.Allergy}} value={{.userinfo.Allergy}} {{end}}
				 maxlength="1200"
				 >
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="MedicalInfo" class="control-label">Informations médicales importantes</label>
				<input type="text" class="form-control" id="MedicalInfo" name="MedicalInfo" placeholder="Information médicale" data-required-error="Veuillez remplir ce champ"
				 {{if .userinfo.MedicalInfo}} value={{.userinfo.MedicalInfo}} {{end}}
				 maxlength="1200"
				 >
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="DriverLicenceVL" class="control-label">Avez-vous le permis de conduire Véhicule Léger?*</label>
				<div> </div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="DriverLicenceVL" value="1"
						{{if eq .userinfo.DriverLicenceVL 1}}
							checked
						{{end}}
					required>Oui</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="DriverLicenceVL" value="2"
					{{if eq .userinfo.DriverLicenceVL 2}}
							checked
					{{end}}
					required>Non</label>
				</div>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="DriverLicencePL" class="control-label">Avez-vous le permis de conduire Poids Lourd?*</label>
				<div> </div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="DriverLicencePL" value="1"
						{{if eq .userinfo.DriverLicencePL 1}}
							checked
						{{end}}
					required>Oui</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="DriverLicencePL" value="2"
					{{if eq .userinfo.DriverLicencePL 2}}
							checked
					{{end}}
					required>Non</label>
				</div>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="FirstAidTraining" class="control-label">Avez-vous suivi une formation aux premiers secours?*</label>
				<div> </div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="FirstAidTraining" value="1"
						{{if eq .userinfo.FirstAidTraining 1}}
							checked
						{{end}}
					required>Oui</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="FirstAidTraining" value="2"
					{{if eq .userinfo.FirstAidTraining 2}}
							checked
					{{end}}
					required>Non</label>
				</div>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="EnglishLevel" class="control-label">Niveau d'anglais*</label>
				<select class="form-control" id="EnglishLevel" name="EnglishLevel" data-required-error="Veuillez indiqué votre niveau d anglais" required>
					<option value="" disabled {{ if eq .userinfo.EnglishLevel 0}} selected {{end}}>Niveau d anglais</option>
					<option value="1" {{ if eq .userinfo.EnglishLevel 1}} selected {{end}}>un peu</option>
					<option value="2" {{ if eq .userinfo.EnglishLevel 2}} selected {{end}}>scolaire</option>
					<option value="3" {{ if eq .userinfo.EnglishLevel 3}} selected {{end}}>bon</option>
					<option value="4" {{ if eq .userinfo.EnglishLevel 4}} selected {{end}}>couramment</option>
					<option value="5" {{ if eq .userinfo.EnglishLevel 5}} selected {{end}}>billingue</option>
				</select>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="OtherLanguage" class="control-label">Quelle(s) autre(s) langue(s) parlez-vous?</label>
				<input type="text" class="form-control" id="OtherLanguage" name="OtherLanguage" placeholder="espagnol/compris, finnois/courament, grunt/billingue"
				 data-required-error="Veuillez remplir ce champ" {{if .userinfo.OtherLanguage}} value={{.userinfo.OtherLanguage}} {{end}}
				 maxlength="1200"
				 >
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>

			<div class="form-group has-feedback">
				<label for="AlreadyBeenBenevolFOS" class="control-label">Si vous avez-déjà été bénévole sur une autre édition du Fall Of Summer, quand et quel poste aviez-vous?</label>
				<input type="text" class="form-control" id="AlreadyBeenBenevolFOS" name="AlreadyBeenBenevolFOS" placeholder="Oui, l'année derniere, j'etais preposé biture au bar VIP. enfin je crois. c est assez flou comme week end."
				 data-required-error="Veuillez remplir ce champ" {{if .userinfo.AlreadyBeenBenevolFOS}} value={{.userinfo.AlreadyBeenBenevolFOS}}
				 {{end}}
				 maxlength="2000"
				 >
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="DidYouCameFOS" class="control-label">Etes-vous déjà venu au Fall Of Summer?*</label>
				<div> </div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="DidYouCameFOS" value="1"
						{{if eq .userinfo.DidYouCameFOS 1}}
							checked
						{{end}}
					required>Oui</label>
				</div>
				<div class="radio-inline">
					<label class="radio-inline"><input type="radio" name="DidYouCameFOS" value="2"
					{{if eq .userinfo.DidYouCameFOS 2}}
							checked
					{{end}}
					required>Non</label>
				</div>
				<div class="help-block with-errors"></div>
			</div>

			<div class="form-group has-feedback">
				<label for="WhatYouWantToDo" class="control-label">Que voulez-vous faire pour aider le Fall Of Summer?(4 choix dans l'ordre de préférence)*</label>
				<div class="row is-table-row">
					<div class="col-sm-6"  style="height: 100%">
						<div class="control-label text-center"> Jobs Disponibles (à faire glisser dans la liste de choix) </div>
						<ul id="whatyoucando" class="list-group"  style="background: Gray;height: 420px">
							<li class="list-group-item text-center" value="1" >accueil artiste</li>
							<li class="list-group-item text-center" value="2" >accueil public /contrôle d'accès</li>
							<li class="list-group-item text-center" value="3" >backline / road</li>
							<li class="list-group-item text-center" value="4" >caisse (tickets ou jetons)</li>
							<li class="list-group-item text-center" value="5" >écocup (rinçage gobelets...)</li>
							<li class="list-group-item text-center" value="6" >environnement</li>
							<li class="list-group-item text-center" value="7" >restauration</li>
							<li class="list-group-item text-center" value="8" >merchandising</li>
							<li class="list-group-item text-center" value="9" >montage / démontage</li>
							<li class="list-group-item text-center" value="10" >runs</li>
							<li class="list-group-item text-center" value="11" >bar</li>
						</ul>
					</div>
  					<div class="col-sm-6" style="height: 100%">
						<div class="control-label text-center"> Votre choix(dans l'ordre de preference, le plus important en haut) </div>
						<ol id="whatyouchoosetodo" class="list-group" style="background: gray;height: 420px">
						</ol>
					</div>
				</div>
				<input type="hidden" class="form-control" id="WhatYouWantToDo" name="WhatYouWantToDo" data-validate="true"
				 data-whatyouwantodo data-required-error="Veuillez signaler vos 4 choix"
				 {{if .what_you_want_to_do}} value={{.what_you_want_to_do}} {{end}}
				 required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="OtherJobs" class="control-label">Autre boulot ou information sur vos choix</label>
				<input type="text" class="form-control" id="OtherJobs" name="OtherJobs" placeholder="Je peux faire reaccordeur de guitare?"
				 data-required-error="Veuillez remplir ce champ" {{if .userinfo.OtherJobs}} value={{.userinfo.OtherJobs}} {{end}}
				 maxlength="1200"
				 >
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="WhenCanYouBeThere" class="control-label">Quand pouvez vous venir?*</label>
				<div class="row">
					<div class="col-sm-5 centered-text montage"><label class="control-label" style="width: 100%; text-align:center;"> Montage </label> </div>
					<div class="col-sm-2 centered-text exploitation"><label class="control-label" style="width: 100%; text-align:center;"> Exploitation </label> </div>
					<div class="col-sm-4 centered-text demontage"><label class="control-label" style="width: 100%; text-align:center;"> Démontage </label> </div>
				</div>
				<div class="row">
					<div class="col-sm-1 centered-text montage"><label class="control-label"> Dimanche 3 septembre </label> </div>
					<div class="col-sm-1 centered-text montage"><label class="control-label"> Lundi 4 septembre </label></div>
					<div class="col-sm-1 centered-text montage"><label class="control-label"> Mardi 5 septembre </label></div>
					<div class="col-sm-1 centered-text montage"><label class="control-label"> Mercredi 6 septembre </label></div>
					<div class="col-sm-1 centered-text montage"><label class="control-label"> Jeudi 7 septembre </label></div>
					<div class="col-sm-1 centered-text exploitation"><label class="control-label"> Vendredi 8 septembre </label></div>
					<div class="col-sm-1 centered-text exploitation"><label class="control-label"> Samedi 9 septembre </label></div>
					<div class="col-sm-1 centered-text demontage"><label class="control-label"> Dimanche 10 septembre </label></div>
					<div class="col-sm-1 centered-text demontage"><label class="control-label"> Lundi 11 septembre </label></div>
					<div class="col-sm-1 centered-text demontage"><label class="control-label"> Mardi 12 septembre </label></div>
					<div class="col-sm-1 centered-text demontage"><label class="control-label"> Mercredi 13 septembre </label></div>
				</div>
				<div id="WhenCanYouBeThereChoices" class="row">
					<div class="col-sm-1 centered-text montage"> <input type="checkbox" data-validate="false" value="1"> </div>
					<div class="col-sm-1 centered-text montage"> <input type="checkbox" data-validate="false" value="2"> </div>
					<div class="col-sm-1 centered-text montage"> <input type="checkbox" data-validate="false" value="4"> </div>
					<div class="col-sm-1 centered-text montage"> <input type="checkbox" data-validate="false" value="8"> </div>
					<div class="col-sm-1 centered-text montage"> <input type="checkbox" data-validate="false" value="16"> </div>
					<div class="col-sm-1 centered-text exploitation"> <input type="checkbox" data-validate="false" value="32"> </div>
					<div class="col-sm-1 centered-text exploitation"> <input type="checkbox" data-validate="false" value="64"> </div>
					<div class="col-sm-1 centered-text demontage"> <input type="checkbox" data-validate="false" value="128"> </div>
					<div class="col-sm-1 centered-text demontage"> <input type="checkbox" data-validate="false" value="256"> </div>
					<div class="col-sm-1 centered-text demontage"> <input type="checkbox" data-validate="false" value="512"> </div>
					<div class="col-sm-1 centered-text demontage"> <input type="checkbox" data-validate="false" value="1024"> </div>
				</div>
				<input type="hidden" class="form-control" id="WhenCanYouBeThere" name="WhenCanYouBeThere"
				data-validate="true"
				data-whencanyoubethere
				data-required-error="Veuillez choisir au moins une journée"
				 {{if .userinfo.WhenCanYouBeThere}} value={{.userinfo.WhenCanYouBeThere}} {{end}}
				required>
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="AlreadyBeenBenevol" class="control-label">Si vous avez-déjà été bénévole sur un Festival, lequel et qu'avez-vous fait?</label>
				<input type="text" class="form-control" id="AlreadyBeenBenevol" name="AlreadyBeenBenevol" placeholder="oui j'ai organisé une garden partie speciale grind et death dans mon jardin. J'ai tiré une rallonge et allumé le barbeuc."
				 data-required-error="Veuillez remplir ce champ" {{if .userinfo.AlreadyBeenBenevol}} value={{.userinfo.AlreadyBeenBenevol}}
				 {{end}}
				 maxlength="1200"
				 >
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<div class="form-group has-feedback">
				<label for="OtherInfo" class="control-label">Voulez vous rajouter quelque chose?</label>
				<input type="text" class="form-control" id="OtherInfo" name="OtherInfo" placeholder="J'vous aime putain!" data-required-error="Veuillez remplir ce champ"
				 {{if .userinfo.OtherInfo}} value={{.userinfo.OtherInfo}} {{end}}
				  maxlength="1200"
				 >
				<span class="glyphicon form-control-feedback" aria-hidden="true"></span>
				<div class="help-block with-errors"></div>
			</div>
			<input type="hidden" name="csrf_token" value="{{.csrf_token}}" />
			<div class="form-group">
				<button type="submit" class="btn btn-primary">Submit</button>
			</div>
		</form>
	</div>
</div>
{{end}}







