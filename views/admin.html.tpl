<button type="button" onclick="location.href='admin/csv';" class="btn btn-warning"> GET Excel </button>
<table id="userlist" class="table table-striped table-bordered" cellspacing="0" width="100%">
        <thead>
            <tr>
                <th>Nom</th>
                <th>Prenom</th>
                <th>Email</th>
                <th> PDF </th>
            </tr>
        </thead>
        <tbody>
            {{range .usersinfo}}
                <tr>
                    <td>{{.Lastname}}</td>
                    <td>{{.Firstname}}</td>
                    <td>{{.ID}}</td>
                    <td><a href="admin/pdf/{{.ID}}">Get PDF</a></td>
                </tr>
            {{end}}
        </tbody>
</table>
<script>
    $(document).ready(function() {
    $('#userlist').DataTable(
        {
            'columns'           : [
                null,   // product code
                null,
                null,   // description
                { 'searchable': false }
            ]
        }
    );
} );
</script>