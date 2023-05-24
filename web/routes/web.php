<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\HomeController;

Route::get('/', [HomeController::class, 'index']);
Route::post('/add-data', [HomeController::class, 'add'])->name('add-data');;
