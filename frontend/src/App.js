import React from 'react';
import './components/style.css';
import {BrowserRouter as Router, Route, Switch} 
      from 'react-router-dom'
import DaftarPinjamBuku from './components/daftarPinjamBuku';
import Navbar from './components/Navbar';
import Home from './components/Home';
import CreatePinjamBuku from './components/CreatePinjamBuku';
import DetailPinjam from './components/DetailPinjam';

function App() {
  return (
    <>
    <div>
        <Router>
              <Navbar />
                <div className="konten">
                    <Switch> 
                          <Route path = "/" exact component =
                              {Home}></Route>
                          <Route path = "/inventory" component = 
                              {DaftarPinjamBuku}></Route>
                          <Route path = "/tambahpinjambuku/:id" component = 
                              {CreatePinjamBuku}></Route>
                          <Route path = "/detailpinjambuku/:id" component = 
                              {DetailPinjam}></Route>
                         </Switch>
                </div>
        </Router>
    </div>
    </>
  );
}

export default App;
