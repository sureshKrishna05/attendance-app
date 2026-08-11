import 'package:flutter/material.dart';
import 'package:frontend/theme/app_theme.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:frontend/config/api_config.dart';
import 'package:frontend/screens/student_dashboard.dart';
import 'package:frontend/screens/faculty_dashboard.dart';
import 'package:shared_preferences/shared_preferences.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  bool isStudent = true;
  bool obscurePassword = true;
  final TextEditingController idController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();

  @override
  void dispose() {
    idController.dispose();
    passwordController.dispose();
    super.dispose();
  }

  bool isLoading = false;

  Future<void> _login() async {
    setState(() => isLoading = true);
    
    try {
      final response = await http.post(
        Uri.parse('${ApiConfig.baseUrl}/auth/login'),
        headers: {
          'Content-Type': 'application/json',
          'X-Api-Key': ApiConfig.publishableKey,
        },
        body: jsonEncode({
          'id': idController.text,
          'password': passwordController.text,
          'role': isStudent ? 'student' : 'faculty',
        }),
      );

      setState(() => isLoading = false);

      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        final prefs = await SharedPreferences.getInstance();
        await prefs.setString('jwt_token', data['token']);
        await prefs.setString('user_id', data['user']['id']);
        await prefs.setString('role', data['user']['role']);
        await prefs.setString('user_name', data['user']['name'] ?? 'User');

        if (!mounted) return;
        if (isStudent) {
          Navigator.pushReplacement(
            context,
            MaterialPageRoute(builder: (context) => const StudentDashboard()),
          );
        } else {
          Navigator.pushReplacement(
            context,
            MaterialPageRoute(builder: (context) => const FacultyDashboard()),
          );
        }
      } else {
        if (!mounted) return;
        
        // Debugging prints to find out what Cloudflare/Backend is actually returning
        print("HTTP Status Code: ${response.statusCode}");
        print("HTTP Body: '${response.body}'");

        String errorMsg = 'Unknown error (Status: ${response.statusCode})';
        try {
          if (response.body.isNotEmpty) {
             errorMsg = jsonDecode(response.body)['error'] ?? errorMsg;
          }
        } catch (e) {
          errorMsg = 'Server returned non-JSON response: Status ${response.statusCode}';
        }
        
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Login failed: $errorMsg')),
        );
      }
    } catch (e) {
      setState(() => isLoading = false);
      if (!mounted) return;
      print("Network/Catch Error: $e");
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Network error: $e')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.surface,
      body: SafeArea(
        child: Center(
          child: SingleChildScrollView(
            padding: const EdgeInsets.symmetric(horizontal: 28.0, vertical: 24.0),
            child: ConstrainedBox(
              constraints: const BoxConstraints(maxWidth: 440),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.stretch,
                children: [
                  const SizedBox(height: 16),

                  // Modern Logo Badge with Subtle Glow/Shadow
                  Center(
                    child: Container(
                      height: 96,
                      width: 96,
                      decoration: BoxDecoration(
                        shape: BoxShape.circle,
                        color: AppColors.info.withValues(alpha: 0.08),
                        border: Border.all(
                          color: AppColors.primary.withValues(alpha: 0.2),
                          width: 4,
                        ),
                        boxShadow: [
                          BoxShadow(
                            color: AppColors.primary.withValues(alpha: 0.12),
                            blurRadius: 20,
                            offset: const Offset(0, 8),
                          ),
                        ],
                      ),
                      child: const Icon(
                        Icons.account_balance_rounded,
                        size: 46,
                        color: AppColors.primary,
                      ),
                    ),
                  ),
                  const SizedBox(height: 24),

                  // University Title
                  const Text(
                    'BHARATHIDASAN UNIVERSITY\nTIRUCHIRAPPALLI',
                    textAlign: TextAlign.center,
                    style: TextStyle(
                      color: AppColors.primary,
                      fontWeight: FontWeight.w800,
                      fontSize: 17,
                      letterSpacing: 0.5,
                      height: 1.3,
                    ),
                  ),
                  const SizedBox(height: 6),
                  const Text(
                    'Attendance Management System',
                    textAlign: TextAlign.center,
                    style: TextStyle(
                      color: AppColors.textSecondary,
                      fontSize: 14,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                  const SizedBox(height: 36),

                  const Text(
                    'Login to your account',
                    textAlign: TextAlign.center,
                    style: TextStyle(
                      fontSize: 18,
                      fontWeight: FontWeight.w700,
                      color: AppColors.textPrimary,
                    ),
                  ),
                  const SizedBox(height: 20),

                  // Animated Segmented Role Toggle
                  Container(
                    height: 52,
                    decoration: BoxDecoration(
                      color: AppColors.surfaceLight,
                      borderRadius: BorderRadius.circular(26),
                      border: Border.all(color: AppColors.border),
                    ),
                    padding: const EdgeInsets.all(4),
                    child: LayoutBuilder(
                      builder: (context, constraints) {
                        final tabWidth = (constraints.maxWidth - 8) / 2;
                        return Stack(
                          children: [
                            // Sliding background pill
                            AnimatedAlign(
                              alignment: isStudent
                                  ? Alignment.centerLeft
                                  : Alignment.centerRight,
                              duration: const Duration(milliseconds: 250),
                              curve: Curves.easeInOutCubic,
                              child: Container(
                                width: tabWidth,
                                decoration: BoxDecoration(
                                  color: AppColors.primary,
                                  borderRadius: BorderRadius.circular(22),
                                  boxShadow: [
                                    BoxShadow(
                                      color: AppColors.primary.withValues(
                                        alpha: 0.3,
                                      ),
                                      blurRadius: 8,
                                      offset: const Offset(0, 3),
                                    ),
                                  ],
                                ),
                              ),
                            ),
                            // Touch targets and labels
                            Row(
                              children: [
                                _buildToggleTab(
                                  label: 'Student',
                                  isSelected: isStudent,
                                  onTap: () => setState(() => isStudent = true),
                                ),
                                _buildToggleTab(
                                  label: 'Faculty',
                                  isSelected: !isStudent,
                                  onTap: () => setState(() => isStudent = false),
                                ),
                              ],
                            ),
                          ],
                        );
                      },
                    ),
                  ),
                  const SizedBox(height: 24),

                  // ID / Email Field
                  TextField(
                    controller: idController,
                    keyboardType: TextInputType.emailAddress,
                    textInputAction: TextInputAction.next,
                    style: const TextStyle(
                      fontSize: 15,
                      color: AppColors.textPrimary,
                      fontWeight: FontWeight.w500,
                    ),
                    decoration: _inputDecoration(
                      hintText: 'University ID / Email',
                      prefixIcon: Icons.person_outline_rounded,
                    ),
                  ),
                  const SizedBox(height: 16),

                  // Password Field
                  TextField(
                    controller: passwordController,
                    obscureText: obscurePassword,
                    textInputAction: TextInputAction.done,
                    onSubmitted: (_) => _login(),
                    style: const TextStyle(
                      fontSize: 15,
                      color: AppColors.textPrimary,
                      fontWeight: FontWeight.w500,
                    ),
                    decoration: _inputDecoration(
                      hintText: 'Password',
                      prefixIcon: Icons.lock_outline_rounded,
                      suffixIcon: IconButton(
                        icon: Icon(
                          obscurePassword
                              ? Icons.visibility_off_outlined
                              : Icons.visibility_outlined,
                          color: AppColors.textSecondary,
                          size: 22,
                        ),
                        onPressed: () => setState(
                          () => obscurePassword = !obscurePassword,
                        ),
                      ),
                    ),
                  ),

                  // Forgot Password Link
                  Align(
                    alignment: Alignment.centerRight,
                    child: TextButton(
                      onPressed: () {},
                      style: TextButton.styleFrom(
                        padding: const EdgeInsets.symmetric(
                          horizontal: 4,
                          vertical: 8,
                        ),
                        minimumSize: Size.zero,
                        tapTargetSize: MaterialTapTargetSize.shrinkWrap,
                      ),
                      child: const Text(
                        'Forgot Password?',
                        style: TextStyle(
                          color: AppColors.primary,
                          fontWeight: FontWeight.w600,
                          fontSize: 13.5,
                        ),
                      ),
                    ),
                  ),
                  const SizedBox(height: 24),

                  // High-Emphasis Login Button
                  SizedBox(
                    height: 52,
                    child: ElevatedButton(
                      onPressed: _login,
                      style: ElevatedButton.styleFrom(
                        backgroundColor: AppColors.primary,
                        foregroundColor: AppColors.secondary,
                        elevation: 2,
                        shadowColor: AppColors.primary.withValues(alpha: 0.4),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(16),
                        ),
                      ),
                      child: isLoading
                          ? const SizedBox(
                              height: 24,
                              width: 24,
                              child: CircularProgressIndicator(
                                color: Colors.white,
                                strokeWidth: 2,
                              ),
                            )
                          : const Text(
                              'Login',
                              style: TextStyle(
                                fontSize: 16,
                                fontWeight: FontWeight.w700,
                                letterSpacing: 0.5,
                              ),
                            ),
                    ),
                  ),
                  const SizedBox(height: 48),

                  // Decorative Footer Illustration
                  Center(
                    child: Column(
                      children: [
                        const Icon(
                          Icons.account_balance_outlined,
                          size: 48,
                          color: AppColors.border,
                        ),
                        const SizedBox(height: 8),
                        Text(
                          'Secure University Portal',
                          style: TextStyle(
                            color: AppColors.textSecondary.withValues(alpha: 0.6),
                            fontSize: 12,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }

  // Helper widget for toggle tab labels
  Widget _buildToggleTab({
    required String label,
    required bool isSelected,
    required VoidCallback onTap,
  }) {
    return Expanded(
      child: GestureDetector(
        behavior: HitTestBehavior.opaque,
        onTap: onTap,
        child: Center(
          child: AnimatedDefaultTextStyle(
            duration: const Duration(milliseconds: 200),
            style: TextStyle(
              color: isSelected ? AppColors.secondary : AppColors.textPrimary,
              fontWeight: isSelected ? FontWeight.w700 : FontWeight.w500,
              fontSize: 14.5,
            ),
            child: Text(label),
          ),
        ),
      ),
    );
  }

  // Consistent Input Decoration Helper
  InputDecoration _inputDecoration({
    required String hintText,
    required IconData prefixIcon,
    Widget? suffixIcon,
  }) {
    return InputDecoration(
      hintText: hintText,
      hintStyle: const TextStyle(
        color: AppColors.textSecondary,
        fontWeight: FontWeight.w400,
        fontSize: 14.5,
      ),
      filled: true,
      fillColor: AppColors.surfaceLight,
      // Removed "const" here because prefixIcon is a method parameter
      prefixIcon: Icon(prefixIcon, color: AppColors.textSecondary, size: 22),
      suffixIcon: suffixIcon,
      contentPadding: const EdgeInsets.symmetric(horizontal: 18, vertical: 16),
      enabledBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(14),
        borderSide: const BorderSide(color: AppColors.border),
      ),
      focusedBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(14),
        borderSide: const BorderSide(color: AppColors.primary, width: 1.5),
      ),
    );
  }
}