import React, { useState } from 'react';
import Papa from 'papaparse';
import { Upload } from 'lucide-react';

export default function BulkUpload() {
  const [data, setData] = useState([]);

  const handleFileUpload = (e) => {
    const file = e.target.files[0];
    if (file) {
      Papa.parse(file, {
        complete: (results) => {
          setData(results.data);
        },
        header: true,
      });
    }
  };

  return (
    <div className="card">
      <h2 className="card-title">Bulk Upload Students / Faculty (CSV/XLSX)</h2>
      <p style={{ color: 'var(--text-muted)', marginBottom: '24px' }}>Upload a CSV file containing user details to bulk register them into the database.</p>
      
      <div style={{ border: '2px dashed var(--border)', borderRadius: '12px', padding: '40px', textAlign: 'center', marginBottom: '24px' }}>
        <Upload size={48} color="var(--primary)" style={{ margin: '0 auto 16px', opacity: 0.5 }} />
        <p style={{ fontWeight: 500, marginBottom: '8px' }}>Drag and drop your file here, or click to browse</p>
        <input type="file" accept=".csv" onChange={handleFileUpload} style={{ marginTop: '16px' }} />
      </div>

      {data.length > 0 && (
        <div className="table-container">
          <h3 style={{ marginBottom: '16px', fontSize: '15px' }}>Preview Data ({data.length} rows)</h3>
          <table>
            <thead>
              <tr>
                {Object.keys(data[0] || {}).map((key) => (
                  <th key={key}>{key}</th>
                ))}
              </tr>
            </thead>
            <tbody>
              {data.slice(0, 5).map((row, i) => (
                <tr key={i}>
                  {Object.values(row).map((val, j) => (
                    <td key={j}>{val}</td>
                  ))}
                </tr>
              ))}
            </tbody>
          </table>
          {data.length > 5 && <p style={{ marginTop: '12px', color: 'var(--text-muted)', fontSize: '13px' }}>Showing first 5 rows...</p>}
          
          <button className="btn btn-primary" style={{ marginTop: '24px' }}>Submit to Database</button>
        </div>
      )}
    </div>
  );
}
