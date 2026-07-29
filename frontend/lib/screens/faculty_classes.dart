import 'package:flutter/material.dart';
import 'package:frontend/theme/app_theme.dart';
import 'package:frontend/screens/faculty_take_attendance.dart';

class FacultyClasses extends StatelessWidget {
  const FacultyClasses({super.key});

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 2,
      child: Scaffold(
        appBar: AppBar(
          backgroundColor: Colors.white,
          elevation: 0,
          leading: IconButton(
            icon: const Icon(Icons.arrow_back, color: AppColors.textPrimary),
            onPressed: () => Navigator.pop(context),
          ),
          title: const Text('My Classes', style: TextStyle(color: AppColors.textPrimary, fontSize: 18, fontWeight: FontWeight.bold)),
          bottom: const TabBar(
            labelColor: AppColors.primary,
            unselectedLabelColor: AppColors.textSecondary,
            indicatorColor: AppColors.primary,
            tabs: [
              Tab(text: 'Assigned Classes'),
              Tab(text: 'All Classes'),
            ],
          ),
        ),
        body: ListView(
          padding: const EdgeInsets.all(16.0),
          children: [
            _buildClassCard(context, 'Data Structures (CSE301)', 'Semester 4 - Section A', '42'),
            const SizedBox(height: 16),
            _buildClassCard(context, 'Operating Systems (CSE304)', 'Semester 4 - Section A', '42'),
            const SizedBox(height: 16),
            _buildClassCard(context, 'Computer Networks (CSE303)', 'Semester 4 - Section B', '40'),
            const SizedBox(height: 16),
            _buildClassCard(context, 'Database Management Systems (CSE302)', 'Semester 4 - Section A', '42'),
          ],
        ),
        floatingActionButton: FloatingActionButton(
          onPressed: () {},
          backgroundColor: AppColors.primary,
          child: const Icon(Icons.add, color: Colors.white),
        ),
      ),
    );
  }

  Widget _buildClassCard(BuildContext context, String title, String subtitle, String studentsCount) {
    return InkWell(
      onTap: () {
        Navigator.push(context, MaterialPageRoute(builder: (context) => const FacultyTakeAttendance()));
      },
      child: Container(
        padding: const EdgeInsets.all(16),
        decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(12),
          border: Border.all(color: Colors.grey.shade200),
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(title, style: const TextStyle(fontWeight: FontWeight.bold, color: AppColors.textPrimary)),
                  const SizedBox(height: 4),
                  Text(subtitle, style: const TextStyle(fontSize: 12, color: AppColors.textSecondary)),
                ],
              ),
            ),
            Column(
              children: [
                Text(studentsCount, style: const TextStyle(fontWeight: FontWeight.bold, fontSize: 18, color: AppColors.textPrimary)),
                const Text('Students', style: TextStyle(fontSize: 10, color: AppColors.textSecondary)),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
