<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\View\View;

class HomeController extends Controller
{
    public function index(): View
    {
        $nationality = $this->curlGet(env('ENDPOINT_API', '') . "nationality", []);

        return view('home', [
            "nationality" => $nationality['content']
        ]);
    }

    public function add(Request $request)
    {
        $input = $request->all();

        $cst_name = @$input["cst_name"];
        $cst_dob = @$input["cst_dob"];
        $cst_email = @$input["cst_email"];
        $cst_phone = @$input["cst_phone"];
        $nationality_id = @$input["nationality_id"];

        $fl_name = @$input["fl_name"];
        $fl_relation = @$input["fl_relation"];
        $fl_dob = @$input["fl_dob"];

        $data_cust = [
            "nationality_id" => (int)$nationality_id,
            "cst_name" => $cst_name,
            "cst_dob" => $cst_dob,
            "cst_phone" => $cst_phone,
            "cst_email" => $cst_email
        ];

        $post_cust = $this->curlPost(env('ENDPOINT_API', '') . "customer", $data_cust);

        if ($post_cust['content'] != "") {
            $i = 0;
            foreach ($fl_dob as $a) {
                $data_fl = [
                    "cst_id" => (int)$post_cust['content'],
                    "fl_relation" => $fl_relation[$i],
                    "fl_name" => $fl_name[$i],
                    "fl_dob" => $fl_dob[$i]
                ];

                $post_fl = $this->curlPost(env('ENDPOINT_API', '') . "family-list", $data_fl);

                $i++;
            }
        }
        return response()->json($post_fl["content"]);
    }

    public function curlGet($endpoint, $param = [])
    {
        $client = new \GuzzleHttp\Client();

        $response = $client->request('GET', $endpoint, ['query' => $param]);

        $statusCode = $response->getStatusCode();
        $content = $response->getBody();

        $content = json_decode($response->getBody(), true);

        return [
            "statusCode" => $statusCode,
            "content" => $content
        ];
    }

    public function curlPost($endpoint, $body)
    {
        $client = new \GuzzleHttp\Client([
            'headers' => ['Content-Type' => 'application/json']
        ]);

        $response = $client->post(
            $endpoint,
            ['body' => json_encode($body)]
        );

        $statusCode = $response->getStatusCode();
        $content = $response->getBody()->getContents();

        return [
            "statusCode" => $statusCode,
            "content" => $content
        ];
    }
}
