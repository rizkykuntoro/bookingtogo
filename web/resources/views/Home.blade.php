<!doctype html>
<html lang="en">

<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="csrf-token" content="{{ csrf_token() }}" />

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css" integrity="sha384-xOolHFLEh07PJGoPkLv1IbcEPTNtaed2xpHsD9ESMhqIYd0nLMwNLD69Npy4HI+N" crossorigin="anonymous">

    <!-- <script src="https://cdn.jsdelivr.net/npm/jquery@3.5.1/dist/jquery.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script> -->
    <script type="text/javascript" src="http://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-Fy6S3B9q64WdZWQUiU+q4/2Lc9npb8tCaSX9FK7E8HnRr0Jz8D6OP9dO5Vg3Q9ct" crossorigin="anonymous"></script>

    <title>Hello, world!</title>
    <style>
        .white {
            color: white;
        }
    </style>
</head>

<body class="container">

    <form class="form">
        <p>
            <b>USER</b>
        </p>

        <div class="form-group">
            <label for="inputCst_name">Nama</label>
            <input type="text" class="form-control" id="inputCst_name" name="cst_name" placeholder="Masukan Nama Anda">
        </div>

        <div class="form-group">
            <label for="inputCst_email">Email</label>
            <input type="email" class="form-control" id="inputCst_email" name="cst_email" placeholder="Masukan Email Anda">
        </div>

        <div class="form-group">
            <label for="inputCst_phone">Phone</label>
            <input type="text" class="form-control" id="inputCst_phone" name="cst_phone" placeholder="Masukan Phone Anda">
        </div>

        <div class="form-group">
            <label for="inputCst_dob">Tanggal Lahir</label>
            <input type="date" class="form-control" id="inputCst_dob" name="cst_dob" value="<?= date("Y-m-d") ?>" placeholder="Pilih Tanggal">
        </div>

        <div class="form-group">
            <label for="inputNationality_id">Kewarganegaraan</label>
            <select class="form-control" id="inputNationality_id" name="nationality_id">
                <?php foreach ($nationality as $n) { ?>
                    <option value="<?= $n['nationality_id'] ?>"><?= $n['nationality_name'] ?> - <?= $n['nationality_code'] ?></option>
                <?php } ?>
            </select>
        </div>

        <hr>
        <p>
            <b>KELUARGA</b><a class="float-right" id="add_family">+ Tambah Keluarga</a>
        </p>

        <div class="form-row form-copy">
            <div class="form-group col-md-3">
                <label for="inputFl_name">Nama</label>
                <input type="text" class="form-control" name="fl_name[]" id="inputFl_name" placeholder="Masukan Nama">
            </div>
            <div class="form-group col-md-3">
                <label for="inputFl_relation">Hubungan</label>
                <input type="text" class="form-control" name="fl_relation[]" id="inputFl_relation" placeholder="Masukan Hubungan">
            </div>
            <div class="form-group col-md-3">
                <label for="inputfl_dob">Tanggal Lahir</label>
                <input type="date" class="form-control" id="inputfl_dob" name="fl_dob[]" placeholder="Pilih Tanggal" value="<?= date("Y-m-d") ?>">
            </div>
            <div class="form-group col-md-3">
                <label class="white">Tanggal Lahir</label>
                <button class="btn btn-danger form-control remove-button">Hapus</button>
            </div>
        </div>
        <div class="form-area"></div>
        <button type="submit" class="btn btn-primary btn-submit">Simpan</button>
    </form>
</body>

<script type="text/javascript">
    $(document).ready(function() {
        $('#add_family').click(function() {
            $('.form-area:first').append('<div class="form-row form-copy"><div class="form-group col-md-3">                <label for="inputFl_name">Nama</label>                <input type="text" class="form-control" name="fl_name[]" id="inputFl_name" placeholder="Masukan Nama">            </div>            <div class="form-group col-md-3">                <label for="inputFl_relation">Hubungan</label>                <input type="text" class="form-control" name="fl_relation[]" id="inputFl_relation" placeholder="Masukan Hubungan">            </div>            <div class="form-group col-md-3">                <label for="inputfl_dob">Tanggal Lahir</label>                <input type="date" value=\'<?= date("Y-m-d") ?>\' class="form-control" id="inputfl_dob" name="fl_dob[]" placeholder="Pilih Tanggal">            </div>            <div class="form-group col-md-3">                <label class="white">Tanggal Lahir</label>                <button class="btn btn-danger form-control remove-button">Hapus</button>            </div></div>').insertAfter('.form-area:last');
        });

        $(document).on("click", ".remove-button", function() {
            $(this).parents(".form-copy").remove()
        })

        $.ajaxSetup({
            headers: {
                'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
            }
        });

        $(".btn-submit").click(function(e) {

            e.preventDefault();

            var cst_name = $("input[name=cst_name]").val();
            var cst_email = $("input[name=cst_email]").val();
            var cst_phone = $("input[name=cst_phone]").val();
            var cst_dob = $("input[name=cst_dob]").val();
            var nationality_id = $('#inputNationality_id').find(":selected").val();

            var fl_name = [];
            $('input[name^=fl_name]').each(function() {
                fl_name.push($(this).val());
            });

            var fl_relation = [];
            $('input[name^=fl_relation]').each(function() {
                fl_relation.push($(this).val());
            });

            var fl_dob = [];
            $('input[name^=fl_dob]').each(function() {
                fl_dob.push($(this).val());
            });

            $.ajax({
                type: 'POST',
                url: "{{ route('add-data') }}",
                data: {
                    cst_name: cst_name,
                    cst_email: cst_email,
                    cst_dob: cst_dob,
                    cst_phone: cst_phone,
                    nationality_id: nationality_id,
                    fl_name: fl_name,
                    fl_relation: fl_relation,
                    fl_dob: fl_dob
                },
                success: function(data) {
                    alert(data)
                    $('.form')[0].reset();
                }
            });

        });
    });
</script>

</html>