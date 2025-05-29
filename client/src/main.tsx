import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import './index.css';
import {CartProvider} from "./contexts/CartContext.tsx";
import Layout from "./layouts/Layout.tsx";
import MainPage from "./pages/MainPage.tsx";
import ProductsPage from "./pages/ProductsPage.tsx";
import CartPage from "./pages/CartPage.tsx";

declare global {
  interface Window {
    _env_: {
      VITE_API_URL: string;
    };
  }
}

export const API_BASE_URL = window._env_?.VITE_API_URL || import.meta.env.VITE_API_URL;

createRoot(document.getElementById('root')!).render(
  <StrictMode>
      <CartProvider>
          <Router>
              <Routes>
                  <Route path='/' element={<Layout><MainPage/></Layout>}/>
                  <Route path='/products' element={<Layout><ProductsPage/></Layout>}/>
                  <Route path='/cart' element={<Layout><CartPage/></Layout>}/>
              </Routes>
          </Router>
      </CartProvider>
  </StrictMode>,
)
