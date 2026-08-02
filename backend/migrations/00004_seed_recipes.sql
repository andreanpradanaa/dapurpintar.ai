-- +goose Up
-- Seed 12 popular Indonesian recipes for MVP discover/detail screens.
-- All are public and available. IDs are deterministic for idempotent re-run.

insert into recipes (id, title, summary, servings, prep_time_minutes, cook_time_minutes, ingredients, instructions, is_public, status)
values
(
  'a0000000-0000-0000-0000-000000000001',
  'Nasi Goreng',
  'Nasi goreng klasik dengan kecap manis, bawang putih, dan telur.',
  2, 10, 15,
  '[{"name":"nasi putih","quantity":"2 piring"},{"name":"bawang putih","quantity":"3 siung"},{"name":"kecap manis","quantity":"2 sdm"},{"name":"telur","quantity":"2 butir"},{"name":"minyak goreng","quantity":"2 sdm"},{"name":"garam","quantity":"secukupnya"},{"name":"bawang merah goreng","quantity":"1 sdm"}]',
  '["Haluskan bawang putih","Panaskan minyak, tumis bawang hingga harum","Masukkan nasi, aduk rata","Tambahkan kecap manis dan garam","Goreng telur mata sapi terpisah","Sajikan nasi goreng dengan telur dan bawang goreng"]',
  true, 'available'
),
(
  'a0000000-0000-0000-0000-000000000002',
  'Soto Ayam',
  'Soto ayam kuning dengan kuah rempah, suwiran ayam, tauge, dan telur rebus.',
  4, 20, 40,
  '[{"name":"ayam","quantity":"500 gr"},{"name":"bawang putih","quantity":"5 siung"},{"name":"bawang merah","quantity":"6 butir"},{"name":"kunyit","quantity":"3 cm"},{"name":"jahe","quantity":"2 cm"},{"name":"serai","quantity":"2 batang"},{"name":"daun jeruk","quantity":"3 lembar"},{"name":"tauge","quantity":"100 gr"},{"name":"telur rebus","quantity":"4 butir"},{"name":"daun bawang","quantity":"1 batang"}]',
  '["Rebus ayam hingga matang, suwir dagingnya","Haluskan bawang merah, bawang putih, kunyit, jahe","Tumis bumbu halus dengan serai dan daun jeruk","Masukkan tumisan ke air kaldu ayam, didihkan","Sajikan dengan tauge, telur rebus, daun bawang, dan nasi"]',
  true, 'available'
),
(
  'a0000000-0000-0000-0000-000000000003',
  'Gado-Gado',
  'Gado-gado Betawi dengan sayuran rebus, tahu, tempe, dan saus kacang.',
  3, 20, 15,
  '[{"name":"tahu","quantity":"200 gr"},{"name":"tempe","quantity":"200 gr"},{"name":"kangkung","quantity":"1 ikat"},{"name":"tauge","quantity":"100 gr"},{"name":"kol","quantity":"100 gr"},{"name":"kentang","quantity":"2 buah"},{"name":"kacang tanah","quantity":"200 gr"},{"name":"gula merah","quantity":"50 gr"},{"name":"cabai","quantity":"3 buah"},{"name":"air asam jawa","quantity":"1 sdm"}]',
  '["Goreng tahu dan tempe hingga kecokelatan","Rebus sayuran satu per satu","Goreng kacang tanah, haluskan dengan cabai dan gula merah","Tambahkan air asam jawa, aduk hingga kental","Tata sayuran, tahu, tempe di piring, siram saus kacang"]',
  true, 'available'
),
(
  'a0000000-0000-0000-0000-000000000004',
  'Rendang',
  'Rendang daging sapi khas Padang yang dimasak perlahan dengan santan dan rempah.',
  6, 30, 180,
  '[{"name":"daging sapi","quantity":"1 kg"},{"name":"santan kental","quantity":"1 liter"},{"name":"bawang merah","quantity":"15 butir"},{"name":"bawang putih","quantity":"8 siung"},{"name":"cabai merah","quantity":"10 buah"},{"name":"lengkuas","quantity":"3 cm"},{"name":"jahe","quantity":"2 cm"},{"name":"serai","quantity":"3 batang"},{"name":"daun kunyit","quantity":"1 lembar"},{"name":"daun jeruk","quantity":"5 lembar"}]',
  '["Potong daging sapi","Haluskan bawang dan cabai","Tumis bumbu halus dengan serai, daun jeruk, daun kunyit","Masukkan daging, aduk hingga berubah warna","Tuang santan, masak dengan api kecil","Aduk sesekali hingga santan mengering dan bumbu meresap (3-4 jam)"]',
  true, 'available'
),
(
  'a0000000-0000-0000-0000-000000000005',
  'Sate Ayam',
  'Sate ayam bakar dengan bumbu kacang dan kecap manis.',
  4, 30, 15,
  '[{"name":"dada ayam","quantity":"500 gr"},{"name":"kecap manis","quantity":"4 sdm"},{"name":"kacang tanah","quantity":"150 gr"},{"name":"bawang putih","quantity":"3 siung"},{"name":"bawang merah","quantity":"5 butir"},{"name":"cabai","quantity":"3 buah"},{"name":"air jeruk limau","quantity":"1 sdm"},{"name":"tusuk sate","quantity":"20 batang"}]',
  '["Potong ayam dadu 2cm","Rendam dengan kecap manis 30 menit","Tusuk ayam ke tusuk sate","Bakar di atas arang sambil dioles bumbu hingga matang","Goreng kacang, haluskan dengan bawang, cabai, dan air","Sajikan sate dengan saus kacang dan jeruk limau"]',
  true, 'available'
),
(
  'a0000000-0000-0000-0000-000000000006',
  'Sayur Asem',
  'Sayur asem segar khas Sunda dengan asam jawa, jagung, dan kacang panjang.',
  4, 15, 25,
  '[{"name":"kacang panjang","quantity":"100 gr"},{"name":"jagung","quantity":"1 buah"},{"name":"labu siam","quantity":"1 buah"},{"name":"melinjo","quantity":"50 gr"},{"name":"asam jawa","quantity":"2 sdm"},{"name":"bawang merah","quantity":"4 butir"},{"name":"bawang putih","quantity":"2 siung"},{"name":"lengkuas","quantity":"2 cm"},{"name":"daun salam","quantity":"2 lembar"},{"name":"gula merah","quantity":"30 gr"}]',
  '["Potong semua sayuran","Rebus air, masukkan jagung dan melinjo","Haluskan bawang, tumis sebentar","Masukkan bumbu dan semua sayuran ke rebusan","Tambahkan asam jawa dan gula merah","Masak hingga sayuran matang, koreksi rasa"]',
  true, 'available'
),
(
  'a0000000-0000-0000-0000-000000000007',
  'Bakso',
  'Bakso sapi kenyal dengan kuah kaldu gurih, mi, dan pangsit.',
  4, 30, 45,
  '[{"name":"daging sapi giling","quantity":"500 gr"},{"name":"tepung tapioka","quantity":"100 gr"},{"name":"putih telur","quantity":"1 butir"},{"name":"bawang putih","quantity":"5 siung"},{"name":"tulang sapi","quantity":"300 gr"},{"name":"mi kuning","quantity":"200 gr"},{"name":"daun bawang","quantity":"2 batang"},{"name":"seledri","quantity":"1 batang"}]',
  '["Haluskan daging, tapioka, putih telur, dan bawang putih","Bentuk bulat-bulat bakso","Rebus tulang untuk kaldu 30 menit","Rebus bakso dalam kaldu hingga mengapung","Rebus mi, sajikan dengan bakso, kuah, dan taburan daun bawang"]',
  true, 'available'
),
(
  'a0000000-0000-0000-0000-000000000008',
  'Mie Goreng',
  'Mie goreng ala rumahan dengan sayuran dan kecap.',
  2, 10, 10,
  '[{"name":"mie telur","quantity":"200 gr"},{"name":"kecap manis","quantity":"3 sdm"},{"name":"bawang putih","quantity":"3 siung"},{"name":"sawi","quantity":"50 gr"},{"name":"wortel","quantity":"1 buah"},{"name":"telur","quantity":"2 butir"},{"name":"minyak goreng","quantity":"2 sdm"},{"name":"garam","quantity":"secukupnya"}]',
  '["Rebus mie hingga al dente, tiriskan","Tumis bawang putih hingga harum","Masukkan telur, orak-arik","Tambahkan sayuran, tumis sebentar","Masukkan mie, kecap manis, garam, aduk rata"]',
  true, 'available'
),
(
  'a0000000-0000-0000-0000-000000000009',
  'Pepes Ikan',
  'Pepes ikan kembung dengan bumbu rempah, daun pisang, dan kemangi.',
  3, 20, 30,
  '[{"name":"ikan kembung","quantity":"500 gr"},{"name":"daun pisang","quantity":"5 lembar"},{"name":"bawang merah","quantity":"6 butir"},{"name":"bawang putih","quantity":"3 siung"},{"name":"kunyit","quantity":"2 cm"},{"name":"kemangi","quantity":"1 ikat"},{"name":"cabai merah","quantity":"5 buah"},{"name":"serai","quantity":"1 batang"},{"name":"garam","quantity":"1 sdt"}]',
  '["Bersihkan ikan, lumuri garam","Haluskan semua bumbu","Campur ikan dengan bumbu halus dan kemangi","Bungkus dalam daun pisang, semat dengan lidi","Kukus 25 menit, lalu bakar sebentar di atas api"]',
  true, 'available'
),
(
  'a0000000-0000-0000-0000-000000000010',
  'Tempe Orek',
  'Tempe orek kering dengan kecap manis dan cabai.',
  2, 10, 15,
  '[{"name":"tempe","quantity":"300 gr"},{"name":"kecap manis","quantity":"3 sdm"},{"name":"bawang merah","quantity":"5 butir"},{"name":"bawang putih","quantity":"3 siung"},{"name":"cabai merah","quantity":"3 buah"},{"name":"lengkuas","quantity":"2 cm"},{"name":"daun salam","quantity":"2 lembar"},{"name":"minyak goreng","quantity":"4 sdm"}]',
  '["Potong tempe tipis, goreng setengah kering","Iris bawang merah, bawang putih, cabai","Tumis bumbu iris hingga harum","Masukkan tempe, kecap manis, daun salam, lengkuas","Aduk hingga bumbu merata dan agak kering"]',
  true, 'available'
),
(
  'a0000000-0000-0000-0000-000000000011',
  'Sambal Goreng Kentang',
  'Kentang goreng dengan sambal, hati ayam, dan petai.',
  3, 15, 25,
  '[{"name":"kentang","quantity":"400 gr"},{"name":"hati ayam","quantity":"150 gr"},{"name":"petai","quantity":"1 papan"},{"name":"cabai merah besar","quantity":"8 buah"},{"name":"bawang merah","quantity":"6 butir"},{"name":"bawang putih","quantity":"3 siung"},{"name":"santan","quantity":"100 ml"},{"name":"daun salam","quantity":"2 lembar"}]',
  '["Kupas dan potong dadu kentang, goreng","Rebus hati ayam, potong dadu","Haluskan cabai dan bawang","Tumis bumbu halus dengan daun salam","Masukkan kentang, hati, petai, santan, masak hingga meresap"]',
  true, 'available'
),
(
  'a0000000-0000-0000-0000-000000000012',
  'Bubur Ayam',
  'Bubur ayam hangat dengan suwiran ayam, cakwe, kacang, dan kecap manis.',
  2, 20, 60,
  '[{"name":"beras","quantity":"150 gr"},{"name":"ayam","quantity":"300 gr"},{"name":"bawang putih","quantity":"4 siung"},{"name":"jahe","quantity":"2 cm"},{"name":"cakwe","quantity":"2 buah"},{"name":"kacang tanah goreng","quantity":"50 gr"},{"name":"kecap manis","quantity":"2 sdm"},{"name":"daun bawang","quantity":"1 batang"}]',
  '["Rebus ayam dengan jahe hingga matang, suwir","Rebus beras dengan banyak air hingga jadi bubur","Tumis bawang putih cincang untuk topping","Sajikan bubur dengan ayam, cakwe, kacang, kecap, dan daun bawang"]',
  true, 'available'
)
on conflict (id) do nothing;

-- +goose Down
delete from recipes where id like 'a0000000-0000-0000-0000-%';
