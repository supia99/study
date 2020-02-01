<?php

$url = "https://www.lantis.jp/imas/release.html";

$html = file_get_contents($url);

// converts all special characters to utf-8
$content = mb_convert_encoding($content, 'HTML-ENTITIES', 'UTF-8');

$dom = new DOMDocument();
$dom->loadHTML($url);

if ($dom->hasChildNodes()) {
    var_dump($dom);
    // foreach(dom->childNodes as childNode) {
    // }
}
