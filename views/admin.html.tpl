<table id="userlist" class="table table-striped table-bordered" cellspacing="0" width="100%">
        <thead>
            <tr>
                <th>Nom</th>
                <th>Prenom</th>
                <th>Email</th>
            </tr>
        </thead>
        <tbody>
            {{range .usersinfo}}
                <tr>
                    <td>{{.Firstname}}</td>
                    <td>{{.Lastname}}</td>
                    <td>{{.ID}}</td>
                </tr>
            {{end}}
        </tbody>
</table>
<script>
    $(document).ready(function() {
    $('#userlist').DataTable();
} );
</script>