import React from 'react';
import './style.css';
import { Link } from 'react-router-dom';

function Home () {
    return (
        <>
        <div className="home">
            <div className="desc">
                <h1>Welcome to <span>Book Loan</span></h1>
                <p>Selamat datang di Book Loan! Jelajahi dunia buku dengan koleksi terbaru dan klasik yang kaya, temukan judul favorit Anda, dan nikmati fitur peminjaman yang nyaman untuk pengalaman membaca yang tak terlupakan.</p>
            </div>
            <button className='btn-home'>
                <Link to="/tambahpinjambuku/_add" style={{ color: 'black', textDecoration: 'none' }}>
                    Isi Data Peminjaman Buku
                </Link>
            </button>
            <button className='btn-home'>
                <Link to="/inventory" style={{ color: 'black', textDecoration: 'none' }}>
                    Lihat Data Peminjaman Buku
                </Link>
            </button>
        </div>
        </>
    );
}

export default Home;
