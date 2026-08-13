import React, { useState } from 'react';
import axios from 'axios';

export default function TimetableEditor() {
  const [formData, setFormData] = useState({
    class_id: '',
    subject_id: '',
    faculty_id: '',
    day: '',
    start_time: '',
    end_time: ''
  });
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState(null);
  const [error, setError] = useState(null);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setMessage(null);
    setError(null);
    try {
      const token = localStorage.getItem('admin_token');
      await axios.post('https://attendance-api.qryptex.in/api/v1/admin/timetable', formData, {
        headers: {
          'X-Admin-Secret': token
        }
      });
      setMessage('Timetable entry created successfully!');
      setFormData({
        class_id: '',
        subject_id: '',
        faculty_id: '',
        day: '',
        start_time: '',
        end_time: ''
      });
    } catch (err) {
      setError(err.response?.data?.message || err.message || 'An error occurred.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="card">
      <h2 className="card-title">Manage Timetable</h2>
      <p style={{ color: 'var(--text-muted)', marginBottom: '24px' }}>Create and modify time slots for classes.</p>
      
      {message && <div style={{ padding: '12px', backgroundColor: '#d4edda', color: '#155724', borderRadius: '4px', marginBottom: '16px' }}>{message}</div>}
      {error && <div style={{ padding: '12px', backgroundColor: '#f8d7da', color: '#721c24', borderRadius: '4px', marginBottom: '16px' }}>{error}</div>}

      <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '16px', maxWidth: '400px' }}>
        <div>
          <label style={{ display: 'block', marginBottom: '8px' }}>Class ID</label>
          <input type="text" name="class_id" value={formData.class_id} onChange={handleChange} required className="input" style={{ width: '100%', padding: '8px', borderRadius: '4px', border: '1px solid var(--border)' }} />
        </div>
        <div>
          <label style={{ display: 'block', marginBottom: '8px' }}>Subject ID</label>
          <input type="text" name="subject_id" value={formData.subject_id} onChange={handleChange} required className="input" style={{ width: '100%', padding: '8px', borderRadius: '4px', border: '1px solid var(--border)' }} />
        </div>
        <div>
          <label style={{ display: 'block', marginBottom: '8px' }}>Faculty ID</label>
          <input type="text" name="faculty_id" value={formData.faculty_id} onChange={handleChange} required className="input" style={{ width: '100%', padding: '8px', borderRadius: '4px', border: '1px solid var(--border)' }} />
        </div>
        <div>
          <label style={{ display: 'block', marginBottom: '8px' }}>Day</label>
          <select name="day" value={formData.day} onChange={handleChange} required className="input" style={{ width: '100%', padding: '8px', borderRadius: '4px', border: '1px solid var(--border)', backgroundColor: 'transparent' }}>
            <option value="">Select Day</option>
            <option value="Monday">Monday</option>
            <option value="Tuesday">Tuesday</option>
            <option value="Wednesday">Wednesday</option>
            <option value="Thursday">Thursday</option>
            <option value="Friday">Friday</option>
            <option value="Saturday">Saturday</option>
            <option value="Sunday">Sunday</option>
          </select>
        </div>
        <div>
          <label style={{ display: 'block', marginBottom: '8px' }}>Start Time</label>
          <input type="time" name="start_time" value={formData.start_time} onChange={handleChange} required className="input" style={{ width: '100%', padding: '8px', borderRadius: '4px', border: '1px solid var(--border)' }} />
        </div>
        <div>
          <label style={{ display: 'block', marginBottom: '8px' }}>End Time</label>
          <input type="time" name="end_time" value={formData.end_time} onChange={handleChange} required className="input" style={{ width: '100%', padding: '8px', borderRadius: '4px', border: '1px solid var(--border)' }} />
        </div>
        <button type="submit" className="btn btn-primary" disabled={loading}>
          {loading ? 'Submitting...' : 'Create Timetable Entry'}
        </button>
      </form>
    </div>
  );
}
