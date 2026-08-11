import 'package:flutter/material.dart';
import 'package:frontend/theme/app_theme.dart';
import 'package:frontend/screens/faculty_classes.dart';
import 'package:frontend/screens/faculty_take_attendance.dart';
import 'package:frontend/screens/faculty_upload_marks.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:frontend/screens/faculty_timetable.dart';

class FacultyDashboard extends StatefulWidget {
  const FacultyDashboard({super.key});

  @override
  State<FacultyDashboard> createState() => _FacultyDashboardState();
}

class _FacultyDashboardState extends State<FacultyDashboard> {
  int _selectedIndex = 0;
  String _userName = 'Faculty';

  @override
  void initState() {
    super.initState();
    _loadUserData();
  }

  Future<void> _loadUserData() async {
    final prefs = await SharedPreferences.getInstance();
    setState(() {
      _userName = prefs.getString('user_name') ?? 'Faculty';
    });
  }

  Widget _buildHomeView() {
    return SingleChildScrollView(
      physics: const BouncingScrollPhysics(),
      padding: const EdgeInsets.symmetric(horizontal: 18.0, vertical: 20.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          // Today's Classes Hero Card
          Container(
            padding: const EdgeInsets.all(22),
            decoration: BoxDecoration(
              gradient: LinearGradient(
                colors: [
                  AppColors.primary,
                  AppColors.primary.withValues(alpha: 0.85),
                ],
                begin: Alignment.topLeft,
                end: Alignment.bottomRight,
              ),
              borderRadius: BorderRadius.circular(20),
              boxShadow: [
                BoxShadow(
                  color: AppColors.primary.withValues(alpha: 0.28),
                  blurRadius: 16,
                  offset: const Offset(0, 8),
                ),
              ],
            ),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'Today\'s Classes',
                      style: TextStyle(
                        color: AppColors.secondary.withValues(alpha: 0.85),
                        fontSize: 14,
                        fontWeight: FontWeight.w500,
                      ),
                    ),
                    const SizedBox(height: 8),
                    const Text(
                      '2',
                      style: TextStyle(
                        color: AppColors.secondary,
                        fontSize: 34,
                        fontWeight: FontWeight.w800,
                        letterSpacing: -0.5,
                      ),
                    ),
                  ],
                ),
                Container(
                  padding: const EdgeInsets.all(16),
                  decoration: BoxDecoration(
                    color: AppColors.secondary.withValues(alpha: 0.12),
                    shape: BoxShape.circle,
                  ),
                  child: const Icon(
                    Icons.calendar_month_rounded,
                    color: AppColors.secondary,
                    size: 42,
                  ),
                ),
              ],
            ),
          ),
          const SizedBox(height: 28),

          // Quick Access Section Header
          const Text(
            'Quick Access',
            style: TextStyle(
              fontSize: 17,
              fontWeight: FontWeight.w700,
              color: AppColors.textPrimary,
              letterSpacing: 0.3,
            ),
          ),
          const SizedBox(height: 16),

          // Quick Access Grid
          GridView.count(
            crossAxisCount: 3,
            shrinkWrap: true,
            physics: const NeverScrollableScrollPhysics(),
            mainAxisSpacing: 14,
            crossAxisSpacing: 14,
            childAspectRatio: 0.95,
            children: [
              _buildQuickAccessIcon(
                Icons.check_circle_outline_rounded,
                'Take Attendance',
                () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) => const FacultyTakeAttendance(),
                    ),
                  );
                },
              ),
              _buildQuickAccessIcon(Icons.class_outlined, 'My Classes', () {
                Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (context) => const FacultyClasses(),
                  ),
                );
              }),
              _buildQuickAccessIcon(
                Icons.upload_file_outlined,
                'Upload Marks',
                () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) => const FacultyUploadMarks(),
                    ),
                  );
                },
              ),
              _buildQuickAccessIcon(Icons.group_outlined, 'Students', () {}),
              _buildQuickAccessIcon(
                Icons.calendar_today_outlined,
                'Timetable',
                () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) => const FacultyTimetable(),
                    ),
                  );
                },
              ),
              _buildQuickAccessIcon(
                Icons.insert_chart_outlined,
                'Reports',
                () {},
              ),
            ],
          ),
          const SizedBox(height: 28),

          // Recent Announcements Section Header
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              const Text(
                'Recent Announcements',
                style: TextStyle(
                  fontSize: 17,
                  fontWeight: FontWeight.w700,
                  color: AppColors.textPrimary,
                  letterSpacing: 0.3,
                ),
              ),
              TextButton(
                onPressed: () {},
                style: TextButton.styleFrom(
                  padding: const EdgeInsets.symmetric(
                    horizontal: 8,
                    vertical: 4,
                  ),
                  minimumSize: Size.zero,
                  tapTargetSize: MaterialTapTargetSize.shrinkWrap,
                ),
                child: const Text(
                  'View All',
                  style: TextStyle(
                    color: AppColors.primary,
                    fontWeight: FontWeight.w600,
                    fontSize: 13.5,
                  ),
                ),
              ),
            ],
          ),
          const SizedBox(height: 12),

          // Announcement Card
          Container(
            padding: const EdgeInsets.all(16),
            decoration: BoxDecoration(
              color: AppColors.surface,
              borderRadius: BorderRadius.circular(16),
              border: Border.all(color: AppColors.border),
              boxShadow: [
                BoxShadow(
                  color: Colors.black.withValues(alpha: 0.02),
                  blurRadius: 8,
                  offset: const Offset(0, 2),
                ),
              ],
            ),
            child: Row(
              children: [
                Container(
                  padding: const EdgeInsets.all(10),
                  decoration: BoxDecoration(
                    color: AppColors.primary.withValues(alpha: 0.08),
                    borderRadius: BorderRadius.circular(12),
                  ),
                  child: const Icon(
                    Icons.campaign_outlined,
                    color: AppColors.primary,
                    size: 22,
                  ),
                ),
                const SizedBox(width: 14),
                const Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        'Department Meeting',
                        style: TextStyle(
                          fontWeight: FontWeight.w600,
                          fontSize: 14.5,
                          color: AppColors.textPrimary,
                        ),
                      ),
                      SizedBox(height: 4),
                      Text(
                        'CSE Department • 18 May 2025',
                        style: TextStyle(
                          fontSize: 12,
                          fontWeight: FontWeight.w500,
                          color: AppColors.textSecondary,
                        ),
                      ),
                    ],
                  ),
                ),
                Container(
                  width: 8,
                  height: 8,
                  decoration: const BoxDecoration(
                    color: AppColors.error,
                    shape: BoxShape.circle,
                  ),
                ),
              ],
            ),
          ),
          const SizedBox(height: 12),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.background,
      appBar: AppBar(
        backgroundColor: AppColors.surface,
        elevation: 0,
        scrolledUnderElevation: 1,
        shadowColor: Colors.black.withValues(alpha: 0.05),
        leading: Padding(
          padding: const EdgeInsets.only(left: 12.0, top: 6.0, bottom: 6.0),
          child: Container(
            decoration: BoxDecoration(
              shape: BoxShape.circle,
              color: AppColors.info.withValues(alpha: 0.08),
              border: Border.all(
                color: AppColors.primary.withValues(alpha: 0.2),
                width: 1.5,
              ),
            ),
            child: const Icon(
              Icons.person_rounded,
              color: AppColors.primary,
              size: 22,
            ),
          ),
        ),
        title: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Text(
              'Good Morning,',
              style: TextStyle(
                fontSize: 12,
                fontWeight: FontWeight.w500,
                color: AppColors.textSecondary,
              ),
            ),
            Text(
              _userName,
              style: const TextStyle(
                fontSize: 16,
                fontWeight: FontWeight.w700,
                color: AppColors.textPrimary,
                letterSpacing: 0.2,
              ),
            ),
            const Text(
              'Assistant Professor - CSE',
              style: TextStyle(
                fontSize: 11.5,
                fontWeight: FontWeight.w500,
                color: AppColors.textSecondary,
              ),
            ),
          ],
        ),
        actions: [
          Padding(
            padding: const EdgeInsets.only(right: 8.0),
            child: IconButton(
              icon: const Icon(
                Icons.notifications_outlined,
                color: AppColors.textPrimary,
                size: 24,
              ),
              onPressed: () {},
            ),
          ),
        ],
      ),
      body: IndexedStack(
        index: _selectedIndex,
        children: [
          _buildHomeView(),
          const FacultyClasses(),
          const FacultyTimetable(),
          const Center(child: Text('Notifications Coming Soon')),
          const Center(child: Text('Profile Settings Coming Soon')),
        ],
      ),
      bottomNavigationBar: Container(
        decoration: BoxDecoration(
          boxShadow: [
            BoxShadow(
              color: Colors.black.withValues(alpha: 0.05),
              blurRadius: 10,
              offset: const Offset(0, -4),
            ),
          ],
        ),
        child: BottomNavigationBar(
          currentIndex: _selectedIndex,
          onTap: (index) => setState(() => _selectedIndex = index),
          type: BottomNavigationBarType.fixed,
          backgroundColor: AppColors.surface,
          selectedItemColor: AppColors.primary,
          unselectedItemColor: AppColors.textSecondary,
          selectedLabelStyle: const TextStyle(
            fontWeight: FontWeight.w600,
            fontSize: 11,
          ),
          unselectedLabelStyle: const TextStyle(
            fontWeight: FontWeight.w500,
            fontSize: 11,
          ),
          showUnselectedLabels: true,
          elevation: 0,
          items: const [
            BottomNavigationBarItem(
              icon: Icon(Icons.home_outlined),
              activeIcon: Icon(Icons.home_rounded),
              label: 'Home',
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.class_outlined),
              activeIcon: Icon(Icons.class_rounded),
              label: 'Classes',
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.group_outlined),
              activeIcon: Icon(Icons.group_rounded),
              label: 'Students',
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.insert_chart_outlined),
              activeIcon: Icon(Icons.insert_chart_rounded),
              label: 'Reports',
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.person_outline),
              activeIcon: Icon(Icons.person_rounded),
              label: 'Profile',
            ),
          ],
        ),
      ),
    );
  }

  // Polished Quick Access Tile
  Widget _buildQuickAccessIcon(
    IconData icon,
    String label,
    VoidCallback onTap,
  ) {
    return Material(
      color: AppColors.surface,
      borderRadius: BorderRadius.circular(16),
      child: InkWell(
        onTap: onTap,
        borderRadius: BorderRadius.circular(16),
        child: Container(
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(16),
            border: Border.all(color: AppColors.border),
            boxShadow: [
              BoxShadow(
                color: Colors.black.withValues(alpha: 0.015),
                blurRadius: 6,
                offset: const Offset(0, 2),
              ),
            ],
          ),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Container(
                padding: const EdgeInsets.all(10),
                decoration: BoxDecoration(
                  color: AppColors.primary.withValues(alpha: 0.08),
                  shape: BoxShape.circle,
                ),
                child: Icon(icon, color: AppColors.primary, size: 26),
              ),
              const SizedBox(height: 10),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 4.0),
                child: Text(
                  label,
                  textAlign: TextAlign.center,
                  maxLines: 1,
                  overflow: TextOverflow.ellipsis,
                  style: const TextStyle(
                    fontSize: 12.5,
                    fontWeight: FontWeight.w600,
                    color: AppColors.textPrimary,
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
