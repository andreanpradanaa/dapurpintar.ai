---
document_id: M1-003
title: Problem Statement
owner: Product Manager
status: Draft
version: 1.0.0
last_updated: 2026-08-02
related_documents:
  - product-vision.md
  - product-mission.md
  - target-users.md
  - user-personas.md
  - value-proposition.md
---

# Problem Statement

## Overview

### Purpose

Dokumen ini menjelaskan permasalahan utama yang ingin diselesaikan oleh DapurPintar AI. Problem Statement menjadi dasar dalam menentukan arah pengembangan produk, prioritas fitur, dan keputusan bisnis.

Seluruh fitur yang dikembangkan harus mampu memberikan solusi terhadap masalah yang dijelaskan dalam dokumen ini.

---

## Executive Summary

Memasak merupakan aktivitas sehari-hari bagi jutaan rumah tangga. Namun proses sebelum memasak sering kali lebih sulit dibandingkan memasaknya sendiri.

Pengguna harus menentukan menu, mengecek stok bahan makanan, membuat daftar belanja, mencari resep, menghitung kebutuhan nutrisi, hingga memastikan bahan makanan tidak terbuang.

Saat ini proses tersebut masih dilakukan menggunakan banyak aplikasi yang tidak saling terhubung.

DapurPintar AI hadir sebagai AI Kitchen Assistant yang menyatukan seluruh proses tersebut ke dalam satu platform yang cerdas dan personal.

---

# Problem Background

Perubahan gaya hidup modern menyebabkan semakin banyak keluarga mengalami kesulitan dalam mengelola aktivitas dapur.

Beberapa faktor yang memengaruhi antara lain:

- Waktu memasak yang semakin terbatas.
- Kesibukan pekerjaan.
- Kurangnya perencanaan makanan.
- Pembelian bahan makanan secara impulsif.
- Sulit menjaga pola makan sehat.
- Banyaknya bahan makanan yang terbuang.

Walaupun informasi resep sangat mudah ditemukan di internet, pengguna tetap harus melakukan banyak pekerjaan secara manual sebelum dapat mulai memasak.

---

# Current Problems

## 1. Daily Meal Decision Fatigue

### Description

Banyak orang menghabiskan waktu setiap hari hanya untuk menjawab pertanyaan sederhana:

> "Hari ini masak apa?"

Keputusan ini diulang setiap hari dan sering menimbulkan kebingungan.

### Impact

- Waktu terbuang.
- Stres menentukan menu.
- Menu menjadi monoton.

---

## 2. Pantry Visibility

### Description

Pengguna tidak mengetahui bahan makanan apa yang masih tersedia di rumah.

Akibatnya mereka:

- membeli bahan yang sudah ada,
- lupa menggunakan bahan yang hampir kedaluwarsa,
- atau tidak sadar bahwa mereka sebenarnya sudah memiliki semua bahan untuk memasak.

### Impact

- Belanja berlebihan.
- Food waste meningkat.
- Pengeluaran rumah tangga bertambah.

---

## 3. Fragmented User Experience

Saat ini pengguna menggunakan banyak aplikasi berbeda.

Contoh:

- Google untuk mencari resep.
- YouTube untuk melihat tutorial.
- Catatan untuk daftar belanja.
- Spreadsheet untuk meal planning.
- ChatGPT untuk bertanya resep.
- Kalender untuk mengatur jadwal.

Semua aplikasi tersebut berdiri sendiri.

Tidak ada platform yang menghubungkan seluruh aktivitas dapur.

---

## 4. Generic Recommendations

Sebagian besar aplikasi hanya memberikan rekomendasi resep populer.

Rekomendasi tersebut tidak mempertimbangkan:

- stok bahan pengguna,
- alergi,
- preferensi makanan,
- tujuan diet,
- anggaran,
- waktu memasak.

Akibatnya rekomendasi sering tidak relevan.

---

## 5. Poor Nutrition Planning

Banyak keluarga ingin hidup lebih sehat tetapi tidak tahu bagaimana menyusun meal plan yang seimbang.

Mereka membutuhkan panduan yang sederhana namun personal.

---

## 6. Food Waste

Sebagian besar rumah tangga pernah mengalami:

- sayur membusuk,
- buah kedaluwarsa,
- bumbu tidak pernah dipakai,
- makanan tersisa yang akhirnya dibuang.

Masalah ini berdampak pada:

- biaya,
- lingkungan,
- efisiensi rumah tangga.

---

# Root Cause Analysis

| Problem | Root Cause |
|----------|------------|
| Bingung memilih menu | Tidak ada rekomendasi yang personal |
| Belanja berlebihan | Tidak mengetahui stok pantry |
| Food waste | Tidak ada pengingat penggunaan bahan |
| Sulit diet | Tidak ada meal planning |
| Banyak aplikasi | Solusi masih terpisah |
| Rekomendasi tidak relevan | AI belum memahami konteks pengguna |

---

# Existing Solutions

Saat ini pengguna mengandalkan berbagai solusi.

| Solution | Strength | Limitation |
|----------|----------|------------|
| Google Search | Banyak informasi | Tidak personal |
| YouTube | Tutorial lengkap | Tidak memahami pantry |
| Aplikasi Resep | Banyak resep | Tidak mengetahui stok bahan |
| ChatGPT | Fleksibel | Tidak menyimpan pantry pengguna |
| Spreadsheet | Fleksibel | Manual |
| Catatan | Mudah | Tidak otomatis |

Belum ada solusi yang menggabungkan seluruh kebutuhan pengguna dalam satu ekosistem.

---

# Opportunity

Perkembangan AI memungkinkan terciptanya platform yang mampu memahami konteks pengguna secara menyeluruh.

AI dapat memanfaatkan informasi seperti:

- stok pantry,
- histori memasak,
- preferensi rasa,
- tujuan nutrisi,
- anggaran,
- waktu memasak,
- jumlah anggota keluarga.

Dengan informasi tersebut, AI dapat memberikan rekomendasi yang jauh lebih relevan dibandingkan solusi yang tersedia saat ini.

---

# Our Problem Statement

> Rumah tangga modern membutuhkan cara yang lebih cerdas untuk mengelola aktivitas memasak sehari-hari. Solusi yang ada masih terfragmentasi, tidak personal, dan belum memanfaatkan AI untuk memahami kondisi nyata pengguna. Akibatnya pengguna mengalami kebingungan menentukan menu, pemborosan bahan makanan, perencanaan belanja yang tidak efisien, dan kesulitan menjaga pola makan sehat.

DapurPintar AI dibangun untuk menyelesaikan masalah tersebut melalui AI Kitchen Assistant yang mampu memahami konteks pengguna dan memberikan rekomendasi yang personal, praktis, dan terintegrasi.

---

# Success Criteria

Masalah dianggap berhasil diselesaikan apabila:

- Pengguna tidak lagi bingung menentukan menu harian.
- Pengguna mengetahui stok pantry secara real-time.
- Food waste berkurang.
- Pengeluaran belanja menjadi lebih efisien.
- Pengguna lebih sering memasak di rumah.
- Meal planning menjadi kebiasaan.
- Pengguna merasa AI benar-benar membantu aktivitas dapur.

---

# Non Goals

Dokumen ini tidak membahas:

- Solusi teknis.
- Arsitektur sistem.
- Database.
- API.
- Teknologi AI.
- UI/UX.
- Strategi implementasi.

Semua akan dijelaskan pada dokumen berikutnya.

---

# Related Documents

- Product Vision
- Product Mission
- Target Users
- User Personas
- Value Proposition
- Product Scope
- Feature Inventory