import { useState, useEffect } from 'react';
import { Routes, Route, Navigate, useNavigate, useLocation } from 'react-router-dom';
import { Users, FileSpreadsheet, Calendar, BookOpen, GraduationCap, LayoutDashboard, LogOut } from 'lucide-react';
import Dashboard from './pages/Dashboard';
import BulkUpload from './pages/BulkUpload';
import ClassAssignments from './pages/ClassAssignments';
import TimetableEditor from './pages/TimetableEditor';

function Login({ onLogin }) {
  const [secretKey, setSecretKey] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    // Simulate backend verification or simple env check
    // Since this is frontend, true secure verification should happen in the backend API
    // But for now, we just mock the login using the local env variable
    if (secretKey === import.meta.env.VITE_ADMIN_SECRET_KEY) {
      onLogin(secretKey);
    } else {
      setError('Invalid Secret Key');
    }
  };

  return (
    <div className="auth-container">
      <div className="auth-card">
        <div className="auth-header">
          <h1>University Admin</h1>
          <p style={{ color: 'var(--text-muted)' }}>Enter your admin secret key to continue</p>
        </div>
        <form onSubmit={handleSubmit}>
          <div className="input-group">
            <label>Admin Secret Key</label>
            <input
              type="password"
              className="input"
              value={secretKey}
              onChange={(e) => setSecretKey(e.target.value)}
              placeholder="••••••••••••"
              required
            />
          </div>
          {error && <p style={{ color: 'var(--danger)', fontSize: '14px', marginBottom: '16px' }}>{error}</p>}
          <button type="submit" className="btn btn-primary" style={{ width: '100%' }}>
            Access Dashboard
          </button>
        </form>
      </div>
    </div>
  );
}

function Layout({ children, onLogout }) {
  const location = useLocation();
  const navItems = [
    { path: '/', label: 'Overview', icon: LayoutDashboard },
    { path: '/bulk-upload', label: 'Bulk Upload', icon: FileSpreadsheet },
    { path: '/assignments', label: 'Class Assignments', icon: Users },
    { path: '/timetable', label: 'Timetable Editor', icon: Calendar },
  ];

  return (
    <div className="app-container">
      <aside className="sidebar">
        <div className="sidebar-header">
          <h2>Admin Portal</h2>
        </div>
        <nav className="sidebar-nav">
          {navItems.map((item) => {
            const Icon = item.icon;
            const isActive = location.pathname === item.path;
            return (
              <a 
                key={item.path} 
                href={item.path} 
                className={`nav-item ${isActive ? 'active' : ''}`}
                onClick={(e) => {
                  e.preventDefault();
                  window.history.pushState({}, '', item.path);
                  window.dispatchEvent(new PopStateEvent('popstate'));
                }}
              >
                <Icon size={20} />
                {item.label}
              </a>
            );
          })}
        </nav>
        <div style={{ padding: '16px', borderTop: '1px solid var(--border)' }}>
          <button onClick={onLogout} className="btn btn-secondary" style={{ width: '100%', justifyContent: 'flex-start' }}>
            <LogOut size={20} />
            Sign Out
          </button>
        </div>
      </aside>
      <main className="main-content">
        <header className="header">
          <h3 style={{ color: 'var(--text-main)', fontSize: '18px' }}>
            {navItems.find(n => n.path === location.pathname)?.label || 'Dashboard'}
          </h3>
          <div style={{ display: 'flex', alignItems: 'center', gap: '12px' }}>
            <div style={{ width: 32, height: 32, borderRadius: '50%', background: 'var(--primary)', color: 'white', display: 'flex', alignItems: 'center', justifyContent: 'center', fontWeight: 'bold' }}>
              A
            </div>
            <span style={{ fontWeight: 500, fontSize: 14 }}>Administrator</span>
          </div>
        </header>
        <div className="page-content">
          {children}
        </div>
      </main>
    </div>
  );
}

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem('admin_token');
    if (token) {
      setIsAuthenticated(true);
    }
  }, []);

  const handleLogin = (key) => {
    localStorage.setItem('admin_token', key);
    setIsAuthenticated(true);
    navigate('/');
  };

  const handleLogout = () => {
    localStorage.removeItem('admin_token');
    setIsAuthenticated(false);
    navigate('/login');
  };

  if (!isAuthenticated) {
    return (
      <Routes>
        <Route path="/login" element={<Login onLogin={handleLogin} />} />
        <Route path="*" element={<Navigate to="/login" replace />} />
      </Routes>
    );
  }

  return (
    <Layout onLogout={handleLogout}>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/bulk-upload" element={<BulkUpload />} />
        <Route path="/assignments" element={<ClassAssignments />} />
        <Route path="/timetable" element={<TimetableEditor />} />
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </Layout>
  );
}

export default App;
