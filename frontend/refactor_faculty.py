import re

with open('lib/screens/faculty_dashboard.dart', 'r') as f:
    content = f.read()

# 1. Imports
content = content.replace("import 'package:frontend/screens/faculty_upload_marks.dart';", "import 'package:frontend/screens/faculty_upload_marks.dart';\nimport 'package:shared_preferences/shared_preferences.dart';")

# 2. State vars
state_vars = """
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
"""
content = re.sub(r'  int _selectedIndex = 0;', state_vars, content)

# 3. Dynamic name
content = content.replace("'Dr. Ramesh',", "_userName,")
content = content.replace("style: TextStyle(", "style: const TextStyle(", 1)
content = content.replace("title: const Column(", "title: Column(")
content = content.replace("children: [", "children: const [", 1)

# 4. Extract body to _buildHomeView and use IndexedStack
body_match = re.search(r'      body: (SingleChildScrollView\(.*?      \),)\n      bottomNavigationBar:', content, re.DOTALL)
if body_match:
    old_body = body_match.group(1)
    # Convert trailing comma to semicolon for return statement
    old_body_code = old_body.rstrip()
    if old_body_code.endswith(','):
        old_body_code = old_body_code[:-1] + ';'
    
    new_body = """      body: IndexedStack(
        index: _selectedIndex,
        children: [
          _buildHomeView(),
          const FacultyClasses(),
          const FacultyTimetable(),
          const Center(child: Text('Notifications Coming Soon')),
          const Center(child: Text('Profile Settings Coming Soon')),
        ],
      ),"""
      
    content = content.replace(f"      body: {old_body}\n      bottomNavigationBar:", f"{new_body}\n      bottomNavigationBar:")
    
    # insert _buildHomeView before build method
    build_home = f"\n  Widget _buildHomeView() {{\n    return {old_body_code}\n  }}\n\n  @override\n  Widget build(BuildContext context)"
    content = content.replace("  @override\n  Widget build(BuildContext context)", build_home)

with open('lib/screens/faculty_dashboard.dart', 'w') as f:
    f.write(content)
