import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider} from "react-router-dom";
import App from './App.jsx'
import Home from './components/Home.jsx';
import About from './components/About.jsx';
import Post from './components/Post.jsx';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        index: true,
        element: <Home />,
      },
      {
        path: "/about",
        element: <About />
      },
      {
        path: "/blog/:id",
        element: <Post />
      },
    ]
  }
])

const root = createRoot(document.getElementById('root'));
root.render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
)