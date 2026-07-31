import 'package:flutter/material.dart';

class AppColors {
  // Brand Core
  static const Color primary = Color(0xFF0D47A1); // University Blue
  static const Color secondary = Color(0xFFFFFFFF); // White
  static const Color accent = Color(0xFFD4AF37); // Minimal Gold

  // Backgrounds & Surfaces
  static const Color background = Color(0xFFF8F9FA); // Screen Background
  static const Color surface = Color(0xFFFFFFFF); // Card & Modal Surface
  static const Color surfaceLight = Color(0xFFF1F4F9); // Tinted Containers

  // Typography
  static const Color textPrimary = Color(0xFF1E293B); // High-contrast slate
  static const Color textSecondary = Color(0xFF64748B); // Muted slate

  // Borders & Dividers
  static const Color border = Color(0xFFE2E8F0);
  static const Color divider = Color(0xFFF1F5F9);

  // Semantic / Status Colors (Useful for attendance & grades)
  static const Color success = Color(0xFF16A34A); // Green (Present / Good)
  static const Color error = Color(0xFFDC2626); // Red (Absent / Alert)
  static const Color warning = Color(0xFFD97706); // Orange
  static const Color info = Color(0xFF2563EB); // Blue
}

final ThemeData appTheme = ThemeData(
  useMaterial3: true,
  primaryColor: AppColors.primary,
  scaffoldBackgroundColor: AppColors.background,
  colorScheme: const ColorScheme.light(
    primary: AppColors.primary,
    onPrimary: AppColors.secondary,
    secondary: AppColors.accent,
    onSecondary: AppColors.textPrimary,
    surface: AppColors.surface,
    onSurface: AppColors.textPrimary,
    error: AppColors.error,
    onError: AppColors.secondary,
  ),

  // Default AppBar Theme
  appBarTheme: AppBarTheme(
    backgroundColor: AppColors.surface,
    foregroundColor: AppColors.textPrimary,
    elevation: 0,
    scrolledUnderElevation: 1,
    shadowColor: Colors.black.withValues(alpha: 0.05),
    centerTitle: false,
    iconTheme: const IconThemeData(color: AppColors.textPrimary),
    titleTextStyle: const TextStyle(
      color: AppColors.textPrimary,
      fontSize: 18,
      fontWeight: FontWeight.w700,
      letterSpacing: 0.2,
    ),
  ),

  // Elevated Button Theme
  elevatedButtonTheme: ElevatedButtonThemeData(
    style: ElevatedButton.styleFrom(
      backgroundColor: AppColors.primary,
      foregroundColor: AppColors.secondary,
      elevation: 2,
      shadowColor: AppColors.primary.withValues(alpha: 0.35),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(16),
      ),
      padding: const EdgeInsets.symmetric(vertical: 16, horizontal: 24),
      textStyle: const TextStyle(
        fontSize: 16,
        fontWeight: FontWeight.w700,
        letterSpacing: 0.4,
      ),
    ),
  ),

  // Text Button Theme
  textButtonTheme: TextButtonThemeData(
    style: TextButton.styleFrom(
      foregroundColor: AppColors.primary,
      textStyle: const TextStyle(
        fontWeight: FontWeight.w600,
        fontSize: 14,
      ),
    ),
  ),

  // Card Theme (Fixed: CardTheme -> CardThemeData)
  cardTheme: CardThemeData(
    color: AppColors.surface,
    elevation: 0,
    margin: EdgeInsets.zero,
    shape: RoundedRectangleBorder(
      borderRadius: BorderRadius.circular(16),
      side: const BorderSide(color: AppColors.border, width: 1),
    ),
  ),

  // Input Field Theme
  inputDecorationTheme: InputDecorationTheme(
    filled: true,
    fillColor: Colors.grey.shade50,
    contentPadding: const EdgeInsets.symmetric(horizontal: 18, vertical: 16),
    hintStyle: const TextStyle(
      color: AppColors.textSecondary,
      fontWeight: FontWeight.w400,
      fontSize: 14.5,
    ),
    border: OutlineInputBorder(
      borderRadius: BorderRadius.circular(14),
      borderSide: const BorderSide(color: AppColors.border),
    ),
    enabledBorder: OutlineInputBorder(
      borderRadius: BorderRadius.circular(14),
      borderSide: const BorderSide(color: AppColors.border),
    ),
    focusedBorder: OutlineInputBorder(
      borderRadius: BorderRadius.circular(14),
      borderSide: const BorderSide(color: AppColors.primary, width: 1.5),
    ),
    errorBorder: OutlineInputBorder(
      borderRadius: BorderRadius.circular(14),
      borderSide: const BorderSide(color: AppColors.error, width: 1.5),
    ),
  ),

  // Divider Theme
  dividerTheme: const DividerThemeData(
    color: AppColors.divider,
    space: 1,
    thickness: 1,
  ),
);