import React from 'react';
import { Users, GraduationCap, BookOpen, Clock } from 'lucide-react';

export default function Dashboard() {
  const stats = [
    { label: 'Total Students', value: '1,245', icon: GraduationCap, color: 'var(--primary)' },
    { label: 'Total Faculty', value: '84', icon: Users, color: 'var(--success)' },
    { label: 'Active Classes', value: '32', icon: BookOpen, color: 'var(--danger)' },
    { label: 'Today Timetables', value: '18', icon: Clock, color: 'var(--secondary)' },
  ];

  return (
    <div>
      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(4, 1fr)', gap: '24px', marginBottom: '32px' }}>
        {stats.map((stat, i) => {
          const Icon = stat.icon;
          return (
            <div key={i} className="card" style={{ marginBottom: 0 }}>
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start' }}>
                <div>
                  <p style={{ color: 'var(--text-muted)', fontSize: '14px', fontWeight: 500, marginBottom: '8px' }}>{stat.label}</p>
                  <h3 style={{ fontSize: '28px' }}>{stat.value}</h3>
                </div>
                <div style={{ padding: '12px', background: `${stat.color}15`, borderRadius: '12px' }}>
                  <Icon color={stat.color} size={24} />
                </div>
              </div>
            </div>
          );
        })}
      </div>

      <div className="card">
        <h2 className="card-title">System Status</h2>
        <p style={{ color: 'var(--text-muted)' }}>The admin portal is running perfectly. Use the sidebar to navigate and manage university data.</p>
      </div>
    </div>
  );
}
