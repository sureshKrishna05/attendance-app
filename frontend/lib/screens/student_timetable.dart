import 'package:flutter/material.dart';
import 'package:frontend/theme/app_theme.dart';

class StudentTimetable extends StatelessWidget {
  const StudentTimetable({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 0,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back, color: AppColors.textPrimary),
          onPressed: () => Navigator.pop(context),
        ),
        title: const Text('Timetable', style: TextStyle(color: AppColors.textPrimary, fontSize: 18, fontWeight: FontWeight.bold)),
        actions: [
          IconButton(icon: const Icon(Icons.calendar_today_outlined, color: AppColors.textPrimary), onPressed: () {}),
        ],
      ),
      body: Column(
        children: [
          // Days Header
          Container(
            color: Colors.white,
            padding: const EdgeInsets.symmetric(vertical: 16),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                _buildDayBadge('Mon', true),
                _buildDayBadge('Tue', false),
                _buildDayBadge('Wed', false),
                _buildDayBadge('Thu', false),
                _buildDayBadge('Fri', false),
                _buildDayBadge('Sat', false),
              ],
            ),
          ),
          const SizedBox(height: 16),
          const Text('Monday, 20 May 2025', style: TextStyle(fontWeight: FontWeight.bold, color: AppColors.textPrimary)),
          const SizedBox(height: 16),
          Expanded(
            child: ListView(
              padding: const EdgeInsets.all(16.0),
              children: [
                _buildTimetableSlot('09:00 AM - 10:00 AM', 'Data Structures', 'CSE301', 'Room 203', 'Dr. Ramesh', AppColors.primary),
                const SizedBox(height: 16),
                _buildTimetableSlot('10:00 AM - 11:00 AM', 'Database Management Systems', 'CSE302', 'Room 204', 'Prof. Kavitha', Colors.purple),
                const SizedBox(height: 16),
                _buildTimetableSlot('11:15 AM - 12:15 PM', 'Computer Networks', 'CSE303', 'Room 205', 'Prof. Sudhir', Colors.green),
                const SizedBox(height: 16),
                _buildTimetableSlot('01:15 PM - 02:15 PM', 'Operating Systems', 'CSE304', 'Room 206', 'Dr. Ramesh', Colors.orange),
                const SizedBox(height: 16),
                _buildTimetableSlot('02:15 PM - 03:15 PM', 'Discrete Mathematics', 'CSE305', 'Room 207', 'Prof. Anitha', Colors.blue),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildDayBadge(String day, bool isSelected) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
      decoration: BoxDecoration(
        color: isSelected ? AppColors.primary : Colors.transparent,
        borderRadius: BorderRadius.circular(20),
      ),
      child: Text(
        day,
        style: TextStyle(
          color: isSelected ? Colors.white : AppColors.textSecondary,
          fontWeight: isSelected ? FontWeight.bold : FontWeight.normal,
        ),
      ),
    );
  }

  Widget _buildTimetableSlot(String time, String subject, String code, String room, String professor, Color indicatorColor) {
    return Container(
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: Colors.grey.shade200),
      ),
      child: IntrinsicHeight(
        child: Row(
          children: [
            Container(
              width: 4,
              decoration: BoxDecoration(
                color: indicatorColor,
                borderRadius: const BorderRadius.only(topLeft: Radius.circular(12), bottomLeft: Radius.circular(12)),
              ),
            ),
            Expanded(
              child: Padding(
                padding: const EdgeInsets.all(16.0),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(time, style: const TextStyle(fontSize: 12, color: AppColors.textSecondary)),
                    const SizedBox(height: 4),
                    Text('$subject ($code)', style: const TextStyle(fontWeight: FontWeight.bold, color: AppColors.textPrimary)),
                    const SizedBox(height: 4),
                    Text('$room • $professor', style: const TextStyle(fontSize: 12, color: AppColors.textSecondary)),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
