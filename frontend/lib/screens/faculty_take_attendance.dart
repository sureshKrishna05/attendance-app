import 'package:flutter/material.dart';
import 'package:frontend/theme/app_theme.dart';

class FacultyTakeAttendance extends StatefulWidget {
  const FacultyTakeAttendance({super.key});

  @override
  State<FacultyTakeAttendance> createState() => _FacultyTakeAttendanceState();
}

class _FacultyTakeAttendanceState extends State<FacultyTakeAttendance> {
  // Store attendance status for each student index. true = present, false = absent.
  final Map<int, bool> _attendance = {
    0: true,
    1: true,
    2: true,
    3: false,
    4: true,
  };

  @override
  Widget build(BuildContext context) {
    int presentCount = _attendance.values.where((v) => v).length;
    int absentCount = _attendance.values.where((v) => !v).length;
    int totalCount = _attendance.length;

    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 0,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back, color: AppColors.textPrimary),
          onPressed: () => Navigator.pop(context),
        ),
        title: const Text('Take Attendance', style: TextStyle(color: AppColors.textPrimary, fontSize: 18, fontWeight: FontWeight.bold)),
        actions: [
          IconButton(icon: const Icon(Icons.more_vert, color: AppColors.textPrimary), onPressed: () {}),
        ],
      ),
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const Text('Data Structures (CSE301)', style: TextStyle(fontWeight: FontWeight.bold, fontSize: 16, color: AppColors.primary)),
                const SizedBox(height: 4),
                const Text('Semester 4 - Section A', style: TextStyle(color: AppColors.textSecondary, fontSize: 12)),
                const SizedBox(height: 12),
                Row(
                  children: [
                    const Icon(Icons.calendar_today_outlined, size: 14, color: AppColors.textSecondary),
                    const SizedBox(width: 4),
                    const Text('20 May 2025', style: TextStyle(fontSize: 12, color: AppColors.textSecondary)),
                    const SizedBox(width: 16),
                    const Icon(Icons.access_time_outlined, size: 14, color: AppColors.textSecondary),
                    const SizedBox(width: 4),
                    const Text('09:00 AM - 10:00 AM', style: TextStyle(fontSize: 12, color: AppColors.textSecondary)),
                  ],
                ),
                const SizedBox(height: 24),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceAround,
                  children: [
                    _buildStatColumn('Present', presentCount.toString(), Colors.green),
                    _buildStatColumn('Absent', absentCount.toString(), Colors.red),
                    _buildStatColumn('Total', totalCount.toString(), AppColors.textPrimary),
                  ],
                ),
                const SizedBox(height: 16),
                TextField(
                  decoration: InputDecoration(
                    hintText: 'Search student',
                    prefixIcon: const Icon(Icons.search),
                    contentPadding: const EdgeInsets.symmetric(vertical: 0),
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(30),
                      borderSide: BorderSide(color: Colors.grey.shade300),
                    ),
                    enabledBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(30),
                      borderSide: BorderSide(color: Colors.grey.shade300),
                    ),
                    filled: true,
                    fillColor: Colors.white,
                  ),
                ),
              ],
            ),
          ),
          const Divider(height: 1),
          Expanded(
            child: ListView(
              children: [
                _buildStudentRow(0, 'Arun Kumar', '21CS001'),
                const Divider(height: 1),
                _buildStudentRow(1, 'Bala Murugan', '21CS002'),
                const Divider(height: 1),
                _buildStudentRow(2, 'Charan Kumar', '21CS003'),
                const Divider(height: 1),
                _buildStudentRow(3, 'Divya Sri', '21CS004'),
                const Divider(height: 1),
                _buildStudentRow(4, 'Gokul Raj', '21CS005'),
              ],
            ),
          ),
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: ElevatedButton(
              onPressed: () {},
              child: const Text('Submit Attendance', style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold)),
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildStatColumn(String label, String count, Color color) {
    return Column(
      children: [
        Text(label, style: TextStyle(color: color, fontSize: 12)),
        const SizedBox(height: 4),
        Text(count, style: TextStyle(color: color, fontSize: 18, fontWeight: FontWeight.bold)),
      ],
    );
  }

  Widget _buildStudentRow(int index, String name, String roll) {
    bool isPresent = _attendance[index] ?? true;

    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
      child: Row(
        children: [
          CircleAvatar(
            backgroundColor: Colors.blue.shade50,
            child: const Icon(Icons.person, color: AppColors.primary),
          ),
          const SizedBox(width: 16),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(name, style: const TextStyle(fontWeight: FontWeight.bold, color: AppColors.textPrimary)),
                Text(roll, style: const TextStyle(fontSize: 12, color: AppColors.textSecondary)),
              ],
            ),
          ),
          InkWell(
            onTap: () => setState(() => _attendance[index] = true),
            child: Container(
              padding: const EdgeInsets.all(4),
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                color: isPresent ? Colors.green : Colors.grey.shade200,
              ),
              child: Icon(Icons.check, size: 20, color: isPresent ? Colors.white : Colors.grey),
            ),
          ),
          const SizedBox(width: 16),
          InkWell(
            onTap: () => setState(() => _attendance[index] = false),
            child: Container(
              padding: const EdgeInsets.all(4),
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                color: !isPresent ? Colors.red : Colors.grey.shade200,
              ),
              child: Icon(Icons.close, size: 20, color: !isPresent ? Colors.white : Colors.grey),
            ),
          ),
        ],
      ),
    );
  }
}
