   M S B A M L   `   `   `  ����M  ICollisionDetection, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null)clr-namespace:UserControlUserControl  5   	       UserControl.MainWindow   5   	   \ XPresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35R NWindowsBase, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35W SPresentationCore, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35D 9http://schemas.microsoft.com/winfx/2006/xaml/presentation    5   	   R NSystem.Xaml, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089:x,http://schemas.microsoft.com/winfx/2006/xaml     5   	   @d2http://schemas.microsoft.com/expression/blend/2008     5   	   Jmc;http://schemas.openxmlformats.org/markup-compatibility/2006     5   	   +localclr-namespace:UserControl     5   	   -   5          � Title$+  %Tower Defence - Place Tower (sandbox)��5   	   $	��480��67   $
��1000��6D    � WindowStartupLocation$ CenterScreen=�6Q    � 
ResizeMode$ NoResize=�6v    � 	Resources 5!         ArcherTowerTemplate �� 
TargetType��%  &	        �� 5"   
   " ��66   .R�5#      � .+�5$      �� 
�������$��2q�6/   
���   �$��40��6W   $��40��6b   ._�5%      �� 
�������$��20(�63   $
��Bold"�6A   $��Center=�6S   $��Center=�6p   !6�   5&      5'      5(      5)      .��5+      �� -   5         ��Name$	 Map��5+      
�������.+�5.   
   �� -   5        $��lblTime��5.      $��48(�6   $��Center=�6-   $��Center=�6J   �� 53   
    �� Orientation$ 
Horizontal=�6    �� Left$ 10��6.   
 �� Top$	 402��6?   .+�54      �� -   5        	 ��Name$	 NewRedTower��54      
���  ��$��40��68   $��40��6C   $��2��6O   
 K� ZIndex$
 1��6�   �� 55      $��10��6   $y�Hidden=�6#   57      �� 5:   
   -   5        $	 TowerPlacement1��5:      $
 99��6-   
�������$��40��6O   $��40��6Z   $	 184��6f   $ 80��6x   6�   �� 5;   
   -   5        $	 NewRedTowerPlacement1��5;      $y�	Collapsed=�63   $
 98��6J   
���  ��$��40��6g   $��40��6r   $��2��6~   $	 184��6�   $ 80��6�   6�   A� 5<   
   -   5         A�Name$$ NewRedTowerCoverAreaPlacement1��5<      $y�	Collapsed=�6:   $
 97��6Q   $	��120��6c   $	��120��6o   $	 144��6|   $ 40��6�   ���	#3F000000���	#260000FF6�   �� 5?   
   -   5        $	 TowerPlacement2��5?      $
 99��6-   
�������$��40��6O   $��40��6Z   $	 264��6f   $ 80��6x   6�   �� 5@   
   -	   5        $	 NewRedTowerPlacement2��5@      $y�	Collapsed=�63   $
 98��6J   
���  ��$��40��6g   $��40��6r   $��2��6~   $	 264��6�   $ 80��6�   6�   A� 5A   
   -
   5        $$ NewRedTowerCoverAreaPlacement2��5A      $y�	Collapsed=�6:   $
 97��6Q   $	��120��6c   $	��120��6o   $	 224��6|   $ 40��6�   ���	#3F000000���	#260000FF6�   �� 5K   
   -   5        $
	 Mob1��5K      $y�Visible=�6"   $
 98��67   
���   �$��20��6V   $��20��6a   $��2��6m   $	 500��6�   $	 135��6�   5M      5O      